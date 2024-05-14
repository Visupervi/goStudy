/**
 * @Author visupervi
 * @Description   请求封装
 * @Date 10:58 上午 2021/11/12
 * @Param
 * @return
 */
export const request = async (url = "", data = {}, type = "POST", method = "fetch") => {

  if (type.toLocaleLowerCase() === "get") {
    let str = "";
    Object.keys(data).map((item, index) => {
      str += item + '=' + data[item] + '&';
    });
    if (str !== "") {
      str = str.substr(0, str.lastIndexOf("&"));
      url = url + "?" + str;
    }
  }

  if (window.fetch && method === "fetch") {
    let requestConfig = {
      // credentials: 'include',
      method: type,
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      mode: "cors",
      // cache: "force-cache"
    };

    // let data = {"1": 2, "3": 4};
    if (type.toLocaleLowerCase() === "post") {
      Object.defineProperty(requestConfig, 'body', {
        value: JSON.stringify(data)
      })
    }

    try {
      const response = await fetch(url, requestConfig);

      if (response) return await response.json();

    } catch (e) {
      throw new Error(e)
    }
  } else {
    return new Promise(((resolve, reject) => {
      let reqObj = !window.XMLHttpRequest ? new ActiveXObject("Microsoft.XMLHTTP") : new XMLHttpRequest();
      let sendData = "";
      if (type.toLowerCase() === "post") {
        sendData = JSON.stringify(data)
      }
      reqObj.open(type, url, true);
      reqObj.setRequestHeader("Content-type", "application/json");
      reqObj.send(sendData);

      reqObj.onreadystatechange = () => {
        if(reqObj.readyState !== 4) return;
        if (reqObj.status >= 200 && reqObj.status < 400) {
          resolve(JSON.parse(reqObj.response));
        } else {
          reject(reqObj.response)
        }
        // if (reqObj.readyState === 4) {
        //   if (reqObj.status === 200) {
        //     resolve(JSON.parse(reqObj.response))
        //   }
        // } else {
        //   reject(reqObj.response)
        // }
      }
    }))
  }

}



