export class ServiceWrapper {
  private base: BasicService
  private prefix: string

  constructor(service: BasicService, path: string) {
    this.base = service
    this.prefix = path
  }

  get<T extends Response | Response[]>(path: string) {
    return this.base.get<T>(this.prefix + path)
  }

  post<T extends Response | Response[]>(path: string, data: Record<string, any>) {
    return this.base.post<T>(this.prefix + path, data)
  }

  toLink(path: string): string {
    return this.base.toLink(this.prefix + path)
  }
}

// from https://github.com/microsoft/TypeScript/issues/1897#issuecomment-1228063688
/** A parsed JSON value. */
export type json = string | number | boolean | null | json[] | { [key: string]: json }

/** A JSON stringify-able value. */
export type jsonable =
  | string
  | number
  | boolean
  | null
  | undefined
  | jsonable[]
  | { [key: string]: jsonable }
  | { toJSON(): jsonable }

export type Response = Record<string, json>

export class BasicService {
  private api: string

  constructor() {
    this.api = 'http://localhost:8080' + import.meta.env.BASE_URL + '/api/'
  }

  get<T extends Response | Response[]>(path: string) {
    return this.request<T>(path, 'GET')
  }

  post<T extends Response | Response[]>(path: string, data: jsonable) {
    return this.request<T>(path, 'POST', data)
  }

  private async request<T extends Response | Response[]>(
    path: string,
    method: string,
    body?: jsonable
  ): Promise<T> {
    const url = this.toLink(path)
    const response = await fetch(url, {
      method,
      headers: {
        'Accept-Version': '1.0.0'
      },
      body: body != null ? JSON.stringify(body) : undefined
    })
    return response.json()
  }

  toLink(path: string): string {
    return (this.api + path).replace(/([^:]\/)\/+/g, '$1')
  }
}

export const Base = new BasicService()
