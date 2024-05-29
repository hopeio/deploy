export interface Request {
  config: Config;
  repo: Repository;
  build: Build;
}

export interface Config {
  data: string;
}

interface Repository {
  id: number;
  uid: number;
  user_id: number;
  namespace: string;
  name: string;
  slug: string;
  scm: string;
  git_http_url: string;
  git_ssh_url: string;
  link: string;
  default_branch: string;
  private: boolean;
  visibility: string;
  active: boolean;
  config: string;
  trusted: boolean;
  protected: boolean;
  ignore_forks: boolean;
  ignore_pulls: boolean;
  cancel_pulls: boolean;
  timeout: number;
  counter: number;
  synced: number;
  created: number;
  updated: number;
  version: number;
}

interface Build {
  id: number;
  repo_id: number;
  number: number;
  parent: number;
  status: string;
  error: string;
  event: string;
  action: string;
  link: string;
  timestamp: number;
  title: string;
  message: string;
  before: string;
  after: string;
  ref: string;
  source_repo: string;
  source: string;
  target: string;
  author_login: string;
  author_name: string;
  author_email: string;
  author_avatar: string;
  sender: string;
  params: string[][];
  cron: string;
  deploy_to: string;
  deploy_id: number;
  started: number;
  finished: number;
  created: number;
  updated: number;
  version: number;
}
