/**
 * @file pukaToken 浦发个性化签名逻辑
 * @author dingyang
 */

/* global KJUR */
import {Base64} from 'js-base64';
import AsyncStorage from '@react-native-community/async-storage';
// import sha256 from 'js-sha256';
import sha256 from 'crypto-js/sha256';
import JSEncrypt from 'jsencrypt';

import {getDeviceId} from '../deviceInfo';
import {getRandom} from '../mathUtil';
import {GLOBAL_CONSTANTS} from '../../constants/global';
import {ALL_ENV, ASYNC_STORAGE_KEY} from '../../constants/constants';

import {nativeLogMsg} from '../../modules/common/rnToNative';
import packageJson from '../../../package.json';

const {getEnv} = GLOBAL_CONSTANTS;

const SDK_VERSION = `V${packageJson.version}`; // sdk版本号
// RSA公钥传输SK
// eslint-disable-next-line max-len
const UAT_TRANSFER_SK = 'MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkTJ9TSQjDxFDsFz2MJoMJjB+bJDZBrG+K085kI/vYv7WrjkdGlhriWi7PncXbM5XGTKwbDq4skhCzS0b6MJAHbh+4CiiUPd+MyAWjqRpSOotZE2MAqrjM+En806wu7buA3ROA8Ztuz8q/tv1vRi4Kdxj+ZR2777iHCs3NFLJbYR1OyMWfZ+Rx6pR7tSpPlK9TT1s9Y+DIugcMEwkp7srMEXjLkdxiHb5InyaCVjsT7ralDhZbepGcwSHyvnOMwd5NoaeIOb0tnZc62bnGuRq6etrrGWWaf53bP0/I8A5Y9ES5ytu1QQDHAkueDupmGsSRM04XzHPHkdVLFYSD1/+EQIDAQAB';
// eslint-disable-next-line max-len
const PAT_TRANSFER_SK = 'MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlXt1/qkqt4wlZeJBgUtFVDvag4dr9VHGjeKkoN51qkZCFs/5VAXnq4NFHYv11vUGHkaT/fun4nVPwW+rdLsr7OVAhHSKeerw/6KYDTErPguwj8CFvQlWZZxGR0A52RTjeGn/pd3kuooa0/BQfJB6au7YKA0hdSnKe8OkybKG5TMar5Z/q8nOzVm/dF8PMIz2z/OZZDjz0DUUczyW2uzFFFZ/MQiYJlHP2uErJlA29R8WzdOEijhJ6YvuLVYttj2MlE9p+xhNsU/6Feq5etQiBArY+yvX4922NccY63K1tMfGKAlwNTc1z4Di5B6WmibRvPDkCaJ+f5xYpgDxnc85bQIDAQAB';

export const getTransferSk = () => {
    const currentEnv = __DEV__ ? ALL_ENV.DEBUG : getEnv();
    if (currentEnv === ALL_ENV.PAT) {
        // 投产pk
        return PAT_TRANSFER_SK;
    }
    // 其他都用UAT pk
    return UAT_TRANSFER_SK;
};

// 签名算法
const SYSTEM_VERSION = 'X-SPDBCCC-System-Version';
const AUTH_DEVICE_ID = 'X-SPDBCCC-Auth-Deviceid';
const SYSTEM_NONCE = 'X-SPDBCCC-System-Nonce';
const SYSTEM_TIMESTAMP = 'X-SPDBCCC-System-Timestamp';
const DIGEST = 'X-SPDBCCC-Digest';
const ALGORITHM_ENCRYPT = 'X-SPDBCCC-Algorithm-Encrypt';
const USER_ID = 'X-SPDBCCC-Auth-Userid';

export const encrypt = async (token, requestBody = {}, withUid = true) => {
    const start = +new Date();
    const timestamp = `${+new Date()}`; // 获取时间戳
    const deviceId = getDeviceId();

    // console.log("deviceId",deviceId);
    const nonceRandom = `${deviceId}${timestamp}${getRandom(0, 10000)}`;
    const uid = await AsyncStorage.getItem(ASYNC_STORAGE_KEY.USER_ID);

    // 提取requestBody摘要信息
    const requestBodyStr = JSON.stringify(requestBody);
    // 序列化数据
    const serializeData = [
        `${SYSTEM_VERSION}=${SDK_VERSION}`,
        `${AUTH_DEVICE_ID}=${deviceId}`,
        `${SYSTEM_NONCE}=${nonceRandom}`,
        `${SYSTEM_TIMESTAMP}=${timestamp}`,
        ...((withUid && uid) ? [`${USER_ID}=${uid}`] : []),
        `${DIGEST}=${requestBodyStr}`
    ];

    const encryptText = serializeData.join('&');

    // RSA加密
    const jse = new JSEncrypt();
    jse.setPublicKey(token);

    const shaRes = sha256(encryptText).toString();
    const serializeEncrypt = jse.encrypt(shaRes);
    // const encryptDevice = jse.encrypt(deviceId);
    // nativeLogMsg('--deviceId--' + deviceId);
    // nativeLogMsg('--base64DeviceId--' + Base64.encode(deviceId) );
    // nativeLogMsg('--encryptDevice--' + encryptDevice);

    return {
        [SYSTEM_VERSION]: SDK_VERSION,
        // [AUTH_DEVICE_ID]: Base64.encode(deviceId),
        [AUTH_DEVICE_ID]: deviceId,
        [SYSTEM_NONCE]: nonceRandom,
        [SYSTEM_TIMESTAMP]: timestamp,
        ...((withUid && uid) ? {[USER_ID]: uid} : {}),
        [ALGORITHM_ENCRYPT]: serializeEncrypt ? Base64.encode(serializeEncrypt) : ''
    };
};
