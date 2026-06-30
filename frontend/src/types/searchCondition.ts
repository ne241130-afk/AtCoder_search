export interface SearchCondition {
    keyword: string
    tags: string[]
    contestType: string
    minDifficulty: number | null
    maxDifficulty: number | null
}