/**
 * @file customRequest
 * @author dingyang
 */

/* global KJUR */

import _ from 'lodash';

import api from '../../configure/api';
import {GLOBAL_CONSTANTS} from '../../constants/global';

import request from '../Request';
import {getTransferSk, encrypt} from './pukaEncrypt';

import {nativeLogMsg} from '../../modules/common/rnToNative';

const {getErrorHtml} = GLOBAL_CONSTANTS;

const ERROR_CODE = {
    TOKEN_ERROR: 'tokenError' // token-error
};
const MAX_TOKEN_REQUEST_COUNT = 10; // token最大尝试次数
const TOKEN_REQUEST_TIMEOUT = 1000; // token每次轮询请求的时间间隔

const MAX_TOKEN_ASK_COUNT = 600; // 业务请求询问token有效最大次数
const TOKEN_ASK_TIMEOUT = 50; // 当token-requesting时其他业务接口轮询token时间间隔

const TOKEN_VALID_BUFFER_TIME = 5 * 60 * 1000; // 前端提前判断token失效，5分钟

// auth-token
function authTokenFunc() {
    // 记录全局唯一token相关信息
    let tokenInfo = {
        createTimestamp: '', // token创建事件，当前并没有用到
        expireTimeStamp: '', // 记录token过期时间
        token: '' // 核心变量：当前token
    };

    // 直接不经验证 直接返回token
    const getAuthTokenDirectly = () => tokenInfo.token;
    // 更新tokenInfo对象
    const setTokenInfo = (token, expires) => {
        tokenInfo = {
            createTimestamp: +new Date(), // token创建时间
            expireTimeStamp: +expires, // token失效时间
            token
        };
    };

    // 记录token获取中，避免其他api重复发出无用请求
    let tokenRequesting = false;
    // 设置tokenRequesting状态
    const setTokenRequesting = isTokenRequesting => {
        tokenRequesting = isTokenRequesting;
    };
    // 获取tokenRequesting状态
    const getTokenRequesting = () => tokenRequesting;

    // 判断token是否有效
    const isTokenValid = () => {
        // token失效事件 & 当前token
        const {expireTimeStamp, token} = tokenInfo;
        // 当前时间
        const currentTimestamp = +new Date();
        if (token && (currentTimestamp + TOKEN_VALID_BUFFER_TIME) < expireTimeStamp) {
            return true;
        }
        return false;
    };

    // 执行session-sk请求
    const execTokenRequestTask = async (resolve, reject, currentCount) => {
        try {
            const silent = currentCount !== MAX_TOKEN_REQUEST_COUNT;
            const sessionAuth = await getSessionAuth(silent);
            resolve(sessionAuth);
        }
        catch (err) {
            // 如果出现302，不重试，直接reject
            if (err && err.code === '302') {
                reject(err);
            }
            else if (currentCount < MAX_TOKEN_REQUEST_COUNT) {
                setTimeout(() => {
                    currentCount = currentCount + 1;
                    execTokenRequestTask(resolve, reject, currentCount);
                }, TOKEN_REQUEST_TIMEOUT);
            }
            else {
                // 超过MAX_TOKEN_REQUEST_COUNT（10）次，抛出异常
                const error = new Error();
                error.message = 'token请求出错';
                error.code = ERROR_CODE.TOKEN_ERROR;
                reject(error);
            }
        }
    };
    // 尝试获取token请求
    const requestToken = () => {
        return new Promise((resolve, reject) => {
            execTokenRequestTask(resolve, reject, 1);
        });
    };

    // 所有业务请求前置的获取签名所要用的token
    // 如果当前token是有效的，直接从tokenInfo中返回token，如果无效，则要发起请求来获取最新有效token
    const getAuthToken = async () => {
        const {token} = tokenInfo;

        // 如果不存在token || (存在token & token已过期)：即token无效
        if (!isTokenValid()) {
            setTokenRequesting(true);
            try {
                // 注意这里不是只发一次请求，异常情况下最多发10次直到获取token
                const sessionAuth = await requestToken();
                const {sessionSK, expireTimeStamp} = sessionAuth;
                setTokenInfo(sessionSK, expireTimeStamp);
                setTokenRequesting(false);
                return sessionSK;
            }
            catch (error) {
                setTokenRequesting(false);
                return Promise.reject(error);
            }
        }
        return token;
    };

    return {
        isTokenValid,
        setTokenInfo,
        getAuthToken,
        getAuthTokenDirectly,
        getTokenRequesting
    };
}

const {
    setTokenInfo,
    getAuthToken,
    getAuthTokenDirectly,
    getTokenRequesting,
    isTokenValid
} = authTokenFunc();

const cusRequest = {
    async post(apiUrl, payload = {}, options = {}) {

        // nativeLogMsg('--------request-params--------');
        // nativeLogMsg(JSON.stringify(payload));

        const withTransferSK = _.get(payload, 'transferSK');
        // 删除transfer-sdk参数
        if (_.isObject(payload) && payload.transferSK) {
            delete payload.transferSK;
        }

        // 判断当前是否在requesting-token当中
        const isRequestingToken = getTokenRequesting();
        // 说明有前置请求已经触发了requesting-token，此时需要强行挂起当前请求
        // 待前置token请求结束之后，继续放行当前请求
        if (isRequestingToken && !withTransferSK) {
            // 故意延迟执行当前事件
            try {
                await runAsync();
            }
            catch (error) {
                return Promise.reject(error);
            }
        }

        // token请求拦截
        let token;
        try {
            token = withTransferSK || await getAuthToken();
        }
        catch (err) {
            return Promise.reject(err);
        }

        // console.log('-------^^^^^^^^^^^^^^^^^^^^^^^^------------');
        // console.log('--apiUrl--', apiUrl);
        // 核心：加密算法
        const {from: cFrom, withUid: cWithUid} = options;
        const withUid = !(cFrom === 'h5' && cWithUid === false);
        const customHeader = await encrypt(token, payload, withUid);
        const decorateOptions = {
            ...options,
            headers: {
                ...(options.headers || {}),
                ...(customHeader)
            }
        };

        return new Promise((resolve, reject) => {
            request.post(apiUrl, payload, decorateOptions)
                .then(res => {
                    resolve(res);
                })
                .catch(error => {
                    const errorCode = _.get(error, 'response.status');
                    if (!withTransferSK && errorCode === 401 && token === getAuthTokenDirectly()) {
                        // 说明token权限验证未通过
                        setTokenInfo('', '');
                    }
                    reject(error);
                });
        });
    }
};

// auth: 请求获取会话私钥 sessionSk
async function getSessionAuth(silent = false) {
    try {
        const sessionAuth = await cusRequest.post(api.authSession(), {
            transferSK: getTransferSk()
        }, {silent});
        return sessionAuth;
    }
    catch (error) {
        throw error;
    }
}

function execTokenTask(resolve, reject, currentTimeoutCount) {
    setTimeout(function () {
        if (isTokenValid()) {
            resolve('success');
        }
        else {
            // 如果出现302，不重试，直接reject
            if (getErrorHtml()) {
                reject();
            }
            else if (currentTimeoutCount >= MAX_TOKEN_ASK_COUNT) {
                const error = new Error();
                error.message = 'token请求超时';
                error.code = ERROR_CODE.TOKEN_ERROR;
                reject(error);
            }
            else {
                currentTimeoutCount = currentTimeoutCount + 1;
                execTokenTask(resolve, reject, currentTimeoutCount);
            }
        }
    }, TOKEN_ASK_TIMEOUT);
}

// 开启await等待
function runAsync() {
    return new Promise((resolve, reject) => {
        execTokenTask(resolve, reject, 1);
    });
}

export default cusRequest;
