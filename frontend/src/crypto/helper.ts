// https://gist.github.com/deiu/2c3208c89fbc91d23226

export function arrayBufferToBase64String(arrayBuffer: ArrayBuffer): string {
    var byteArray = new Uint8Array(arrayBuffer)
    var byteString = ''
    for (var i = 0; i < byteArray.byteLength; i++) {
        byteString += String.fromCharCode(byteArray[i])
    }
    return btoa(byteString)
}

export function base64StringToArrayBuffer(b64str: string): ArrayBuffer {
    var byteStr = atob(b64str)
    var bytes = new Uint8Array(byteStr.length)
    for (var i = 0; i < byteStr.length; i++) {
        bytes[i] = byteStr.charCodeAt(i)
    }
    return bytes.buffer
}

export function textToArrayBuffer(str: string): ArrayBuffer {
    var buf = unescape(encodeURIComponent(str)) // 2 bytes for each char
    var bufView = new Uint8Array(buf.length)
    for (var i = 0; i < buf.length; i++) {
        bufView[i] = buf.charCodeAt(i)
    }
    return bufView
}

export function arrayBufferToText(arrayBuffer: ArrayBuffer): string {
    var byteArray = new Uint8Array(arrayBuffer)
    var str = ''
    for (var i = 0; i < byteArray.byteLength; i++) {
        str += String.fromCharCode(byteArray[i])
    }
    return str
}