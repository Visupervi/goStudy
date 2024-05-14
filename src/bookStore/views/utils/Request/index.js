/**
 * @file utils/Request.js 针对bce业务做过二次封装Request工具库
 * @author zhangzhe(zhangzhe@baidu.com)
 */

import Ajax, {AjaxOption} from './Ajax';
import _ from 'lodash';

export class Request extends Ajax {
    static instance;

    constructor() {
        super();
        if (Request.instance) {
            return Request.instance;
        }
        Request.instance = this;
    }

    /*
    beforeSend(options) {
        // 可以对data进行操作
        // 可以对header进行操作
    }
    */

    failHandler(response, options) {
        if (_.isObject(options) && options.from  === 'h5') {
            // 如果请求来源于, 则直接把消息返回给h5自行处理
            return response;
        }
        else if (response.message) {
            if (response.message.redirect) {
                window.location.href = response.message.redirect;
            }
            if (options && !options.silent) {
                // response.message.global && message.error(response.message.global);
            }
            return response.message;
        }
        return {};
    }

    beforeSuccess(response, options) {
        if (response.success) {
            // 成功之后的返回
            let result = response.page || response.result;
            if (!_.isUndefined(response.code) && _.isUndefined(_.get(result, 'code'))) {
                result = {
                    ...result,
                    code: response.code
                };
            }
            if (!_.isUndefined(response.message) && _.isUndefined(_.get(result, 'message'))) {
                result = {
                    ...result,
                    message: response.message
                };
            }
            return result;
        }
        return Promise.reject(this.failHandler(response, options));
    }
}

export default new Request();
