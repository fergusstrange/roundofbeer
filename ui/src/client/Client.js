import axios from 'axios'

const defaultClient = axios.create({});

export class ApiClient {
    constructor() {
        this.options = {
            host: 'https://api.roundof.beer',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        }
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