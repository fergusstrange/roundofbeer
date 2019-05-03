import axios from 'axios'

const defaultClient = axios.create({});
const defaultOptions = {
    host: 'api.roundof.beer',
    headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
    }
};

export class ApiClient {
    constructor(options) {
        this.options = Object.assign({}, defaultOptions)
    }

    request = (method, path, data) => {
        return defaultClient({
            url: path,
            baseURL: this.options.host,
            method: method,
            headers: this.options.headers,
            data: data,
            credentials: 'same-origin'
        });
    }
}