export interface Problem {
  id: string
  title: string
  contest: string
  difficulty: number
  tags: string[]
  atCoderUrl: string
  problemsUrl: string
  solved?: boolean
  favorite?: boolean
}

export interface UserProblemStatus {
  problemId: string
  userId: string
  solved: boolean
  favorite: boolean
  memo: string
  updatedAt: number
}