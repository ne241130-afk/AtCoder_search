import axios from "axios"

import type { Problem, Tag, UserProblemStatus } from "../types/problem"
import type { PaginatedProblems } from "../types/paginatedProblems"
import type { SearchCondition } from "../types/searchCondition"

const api = axios.create({
    baseURL:"http://localhost:8081"
})

export async function getTags(): Promise<Tag[]> {
    const res = await api.get("/api/tags")
    return res.data
}

export async function createTag(name: string): Promise<Tag> {
    const res = await api.post("/api/tags", { name })
    return res.data
}

export async function attachProblemTag(problemId: string, tagId: number): Promise<void> {
    await api.post(`/api/problems/${problemId}/tags`, { tag_id: tagId })
}

export async function detachProblemTag(problemId: string, tagId: number): Promise<void> {
    await api.delete(`/api/problems/${problemId}/tags/${tagId}`)
}

export async function getProblemById(id: string): Promise<Problem> {
    const res = await api.get(`/api/problems/${id}`)
    return res.data
}

export async function getProblems(
    condition:SearchCondition,
    page:number,
    limit:number
):Promise<PaginatedProblems>{

    const res=await api.get("/api/problems",{

        params:{

            keyword:condition.keyword,

            tags:condition.tags.join(","),

            contestType:condition.contestType,

            minDifficulty:condition.minDifficulty ?? "",

            maxDifficulty:condition.maxDifficulty ?? "",
            page,
            limit

        }

    })

    return res.data

}

export async function getProblemStatus(problemId: string): Promise<UserProblemStatus> {
    const res = await api.get(`/api/problems/${problemId}/status`)
    return res.data
}

export async function updateProblemSolved(problemId: string, solved: boolean): Promise<void> {
    await api.patch(`/api/problems/${problemId}/solved`, { solved })
}

export async function updateProblemFavorite(problemId: string, favorite: boolean): Promise<void> {
    await api.patch(`/api/problems/${problemId}/favorite`, { favorite })
}

export async function updateProblemMemo(problemId: string, memo: string): Promise<void> {
    await api.patch(`/api/problems/${problemId}/memo`, { memo })
}