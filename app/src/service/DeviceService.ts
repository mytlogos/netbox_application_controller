import { ServiceWrapper, type BasicService } from './BasicService'

export interface Device {
  Host: string
}

export default class DeviceService {
  #base: ServiceWrapper

  constructor(base: BasicService) {
    this.#base = new ServiceWrapper(base, '/agent')
  }

  async getAll() {
    return this.#base.get('/all')
  }

  async get(uuid: string) {
    return this.#base.get('/get/' + uuid)
  }
}
