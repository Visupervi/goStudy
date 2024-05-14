/**
 * @file utils/Request/Ajax.js Ajax库封装
 * @author zhangzhe(zhangzhe@baidu.com)
 */

import React from 'react';
import {Text, View} from 'react-native';
import _ from 'lodash';
import axios from 'axios';
import {Toast, Modal} from '@ant-design/react-native';
// import NetInfo from '@react-native-community/netinfo';

import {GLOBAL_CONSTANTS} from '../../constants/global';
import {px2dp} from '../../modules/common/helper';

const {getErrorHtml, setErrorHtml} = GLOBAL_CONSTANTS;

const DEFAULT_ERROR_MSG = '系统开小差了，请稍候再试';
const NETWORK_ERROR_MSG = '网络不给力，请稍候再试';
const REQUEST_LIMIT_ERROR_MSG = '当前访问人数较多，请稍后再试';
const UPDATE_ERROR_MSG = '请更新至最新版本！';
const UPDATE_ERROR_CODE = '1401';
export const ERROR_LOGIN_MUST_CODE = 403; // 必须要登录code

const codeMessage = {
    400: DEFAULT_ERROR_MSG, // 发出的请求有错误，服务器没有进行新建或修改数据的操作
    401: DEFAULT_ERROR_MSG, // 用户没有权限（令牌、用户名、密码错误）
    403: DEFAULT_ERROR_MSG, // 用户得到授权，但是访问是被禁止的
    404: DEFAULT_ERROR_MSG, // 发出的请求针对的是不存在的记录，服务器没有进行操作
    406: DEFAULT_ERROR_MSG, // 请求的格式不可得
    410: DEFAULT_ERROR_MSG, // 请求的资源被永久删除，且不会再得到的
    422: DEFAULT_ERROR_MSG, // 当创建一个对象时，发生一个验证错误
    429: REQUEST_LIMIT_ERROR_MSG, // 排队请求数超过最大限制数
    500: DEFAULT_ERROR_MSG, // 服务器发生错误，请检查服务器
    502: DEFAULT_ERROR_MSG, // 网关错误
    503: DEFAULT_ERROR_MSG, // 服务不可用，服务器暂时过载或维护
    504: DEFAULT_ERROR_MSG, // 网关超时

    // 自定义code
    [UPDATE_ERROR_CODE]: UPDATE_ERROR_MSG // 版本更新提示信息
};

axios.defaults.withCredentials = true;
axios.defaults.headers.common.Accept = 'application/json';
axios.defaults.headers.common['Content-Type'] = 'application/json';
axios.defaults.timeout = 30000; // 设置超时时30s

const HEADER_PUBLIC_KEY = 'x-spdbccc-algorithm-publickey';
const HEADER_PUBLIC_KEY_EXPIRETIME = 'x-spdbccc-algorithm-publickey-expiretime';

function UpdateElement() {
    return (<Text allowFontScaling={false} style={{
        marginTop: px2dp(10),
        fontSize: px2dp(28),
        color: '#555',
        justifyContent: 'center'
    }}>{UPDATE_ERROR_MSG}</Text>);
}
export default class Ajax {

    setHeaders(headers) {
        axios.defaults.headers.common = {...axios.defaults.headers.common, ...headers};
    }

    beforeSend(options) {
        return options;
    }

    beforeSuccess(data, options) {
        if (options || !options) {
            return data;
        }
    }

    request(url, options) {
        const opts
            = typeof this.beforeSend === 'function'
                ? this.beforeSend(options) || options
                : options;

        return axios({url, ...opts})
            .then(response => {
                const {status, statusText, headers = {}} = response;

                // 针对浦发，在进行token请求事，header返回信息
                const publicKey = headers[HEADER_PUBLIC_KEY];
                const publicKeyExpireTime = headers[HEADER_PUBLIC_KEY_EXPIRETIME];

                if (status > 400 || status < 200) {
                    const error = new Error(
                        `API ${url} status is ${status} (${statusText})`
                    );
                    return Promise.reject(error);
                }

                // 对header信息进行处理，如果有header 系统信息，直接返回【计算session_sk请求】
                // 放到result让下一步可以拿到
                if (publicKey && publicKeyExpireTime) {
                    return {
                        success: true,
                        result: {
                            sessionSK: publicKey,
                            expireTimeStamp: publicKeyExpireTime
                        }
                    };
                }
                // 其他请求
                return response.data;
            })
            .then(data => {
                if (_.isString(data) && /html>/.test(data)) {
                    // 301 or 302
                    if (getErrorHtml() === '') {
                        setErrorHtml(data);
                    }
                    const error = new Error();
                    error.code = '302';
                    return Promise.reject(error);
                }
                if (_.isObject(data) && data.code === UPDATE_ERROR_CODE) {
                    const error = new Error(UPDATE_ERROR_MSG);
                    error.code = UPDATE_ERROR_CODE;
                    error.response = {};
                    error.response.status = UPDATE_ERROR_CODE;
                    return Promise.reject(error);
                }
                if (typeof this.beforeSuccess === 'function') {
                    return this.beforeSuccess(data, options);
                }


                return data;
            })
            .catch(error => {
                console.log('url', url);
                console.log('AjaxError', error);
                // 日志操作接口报错不出提示
                if (/action\/batch/.test(url)) {
                    return;
                }
                if (error && error.message && (!options || (options && !options.silent))) {
                    const errorCode = _.get(error, 'response.status');
                    let errorMessage = (errorCode && codeMessage[errorCode]) ? codeMessage[errorCode] : error.message;
                    // if (/Error/.test(errorMessage) || /timeout/.test(errorMessage)
                    //     || /((429)|(401))$/.test(errorMessage)) {
                    //     errorMessage = DEFAULT_ERROR_MSG;
                    // }
                    if (/Error/.test(errorMessage) || /timeout/.test(errorMessage)) {
                        if (errorMessage === 'Network Error') {
                            Toast.info(NETWORK_ERROR_MSG, 2);
                        }
                        errorMessage = DEFAULT_ERROR_MSG;
                    } else if (errorCode === UPDATE_ERROR_CODE) {
                        Modal.alert(
                            '升级提示',
                            (<UpdateElement />),
                            [{
                                text: '确定',
                                onPress: () => console.log('ok'),
                                style: {color: '#445BCF'}
                            }]
                        );
                    } else if (!!errorMessage && errorCode !== UPDATE_ERROR_CODE) {
                        Toast.info(errorMessage, 2);
                    } else {
                        Toast.info(NETWORK_ERROR_MSG, 2);
                    }
                    // 检查是否是网络问题
                    // NetInfo.fetch().then(state => {
                    //     if (state.isConnected) {
                    //         if (!!errorMessage && errorCode !== UPDATE_ERROR_CODE) {
                    //             Toast.info(errorMessage, 2);
                    //         }
                    //         else if (errorCode === UPDATE_ERROR_CODE) {
                    //             Modal.alert(
                    //                 '升级提示',
                    //                 (<UpdateElement />),
                    //                 [{
                    //                     text: '确定',
                    //                     onPress: () => console.log('ok'),
                    //                     style: {color: '#445BCF'}
                    //                 }]
                    //             );
                    //         }
                    //     }
                    //     else {
                    //         Toast.info(NETWORK_ERROR_MSG, 2);
                    //     }
                    // });
                }
                return Promise.reject(error);
            });
    }

    async post(url, data, options = {}) {
        const opts = {
            method: 'post',
            data,
            ...options
        };

        return this.request(url, opts);
    }

    get(url, data, options) {
        const opts = {
            method: 'get',
            data,
            ...options
        };
        return this.request(url, opts);
    }

    put(url, data, options) {
        const opts = {
            method: 'put',
            data,
            ...options
        };
        return this.request(url, opts);
    }
}
