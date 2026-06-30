import type { Problem } from './problem'

export interface PaginatedProblems {
  items: Problem[]
  total: number
  page: number
  limit: number
}
