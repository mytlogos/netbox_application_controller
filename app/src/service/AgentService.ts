import { ServiceWrapper, type BasicService } from './BasicService'

export interface DeployAgent {
  Host: string
  Type: 'manual' | 'auto'
}

export interface Host {
  Uuid: string
  Agent: Agent
  Device?: {
    Name:            string
    HostId:          string
    Arch:            string
    KernelVersion:   string
    Platform:        string
    PlatformFamily:  string
    PlatformVersion: string
    OS:              string
    Boottime:        number
    Uptime:          number
  }
}

export interface Agent {
  DeploymentDate: Date
  LastUpdate: Date | null
  State: string
  Version: string
}

export default class AgentService {
  private base: ServiceWrapper

  constructor(base: BasicService) {
    this.base = new ServiceWrapper(base, '/user/agent')
  }

  async deploy(config: DeployAgent): Promise<Host> {
    return this.base.post('/deploy', config) as unknown as Host
  }

  async getAll(): Promise<Host[]> {
    return this.base.get('/all') as unknown as Host[]
  }

  getConfig(uuid: string): Promise<Record<string, any>> {
    return this.base.get('/config/' + uuid)
  }
}
