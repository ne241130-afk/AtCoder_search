export interface Tag {
  id: number
  name: string
}

export interface Problem {
  id: string
  title: string
  contest: string
  difficulty: number
  tags: Tag[]
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