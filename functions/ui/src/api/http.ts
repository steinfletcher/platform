
export enum FetchState {
    PENDING = 'PENDING',
    SUCCESS = 'SUCCESS',
    FAILURE = 'FAILURE',
    UNKNOWN = 'UNKNOWN',
}

export enum Method {
    GET = 'GET',
    POST = 'POST',
    DELETE = 'DELETE',
    PUT = 'PUT',
}

export enum ErrorCode {
    JSON_PARSE_ERROR = 'JSON_PARSE_ERROR',
    NETWORK_CLIENT_ERROR = 'NETWORK_CLIENT_ERROR'
}

export const apiUrl = (subdomain: string, path: string): string => {
    const host = window.location.host;
    if (host.includes("localhost")) {
        return window.location.host;
    }
    return `https://${subdomain}.${host}${path}`;
};

/**
 * request returns a promise with the result of the api call. It never
 * throws errors, so users are guaranteed to receive a resolved promise with
 * either data or errors. Therefore users should not use try/catch semantics.
 * Usage:
 *  const {data, error} = await request<>
 */
export const request = <T>(
    method: Method, url: string,
    body: object | undefined = undefined,
): Promise<Result<T>> =>
    errorAdapter(doRequest(method, url, body));

/**
 * Represents the result of an api call. Contains the response {@code data} T, the raw {@code response}
 * and any service {@code errors}
 */
export class Result<T> {
    readonly error: ServiceError;
    readonly data: T;
    readonly response: Response;

    constructor(error: ServiceError, response: Response, data: T) {
        this.error = error;
        this.data = data;
        this.response = response
    }

    isClientError(): boolean {
        if (this.response) {
            return this.response.status >= 400 && this.response.status < 500
        }
        return false
    }

    isUnauthorized(): boolean {
        if (this.response) {
            return this.response.status == 401
        }
        return false
    }

    isSuccess(): boolean {
        if (this.response) {
            return this.response.status >= 200 && this.response.status < 400
        }
        return false
    }

    isServerError(): boolean {
        if (this.response) {
            return this.response.status >= 500
        }
        return true
    }
}

interface ServiceError {
    code: string;
    description: string;
}

type ResponseHandler<T> = (response: Response) => Promise<Result<T>>

const handleResponse = async <T>(resp: Response): Promise<Result<T>> => {
    let responseJSON;
    try {
        const text = await resp.text();
        if (text && text.length > 0) {
            responseJSON = JSON.parse(text)
        }
    } catch (error) {
        const generalClientError: ServiceError = {
            code: ErrorCode.JSON_PARSE_ERROR,
            description: error.toString()
        };
        return new Result<T>(generalClientError, resp, null as any)
    }

    if (resp.status >= 200 && resp.status < 400) {
        return new Result<T>({} as ServiceError, resp, responseJSON)
    } else {
        let apiErrors = responseJSON as ServiceError;
        return new Result<T>(apiErrors, resp, null as any)
    }
};

const doRequest = async <T>(method: Method, url: string, body: any): Promise<Result<T>> => {
    const resp: Response = await fetch(url, {
        method: method,
        credentials: 'include',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
        body: body ? JSON.stringify(body) : null
    });
    const responseHandler: ResponseHandler<T> = (response: Response) => handleResponse(response);
    return responseHandler(resp)
};

/**
 * Handles any errors thrown from {@code promise} and resolves those errors as data
 */
const errorAdapter = <T>(promise: Promise<Result<T>>): Promise<Result<T>> => {
    return promise
        .then(result => result)
        .catch(error => {
            const networkClientError: ServiceError = {
                code: ErrorCode.NETWORK_CLIENT_ERROR,
                description: error.toString()
            };
            return new Result<T>(networkClientError, null as any, null as any)
        })
};

export const encodeQueryParams = (params: { [key: string]: any; }): string => {
    if (!params || !Object.keys(params).length) {
        return ''
    }

    const initialValue: string[] = [];
    const pairs = Object.keys(params).reduce((acc, key) => {
        const value = params[key];
        if (value === undefined) {
            return acc
        }

        if (Array.isArray(value)) {
            value.forEach((v: any) => acc.push(`${key}=${encodeURIComponent(v)}`))
        } else {
            acc.push(`${key}=${encodeURIComponent(value)}`)
        }
        return acc
    }, initialValue);

    if (pairs.length) {
        return `?${pairs.join('&')}`
    }

    return ''
};
