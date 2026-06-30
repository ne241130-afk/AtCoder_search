import axios from "axios"

import type { Problem } from "../types/problem"
import type { SearchCondition } from "../types/searchCondition"

const api = axios.create({
    baseURL:"http://localhost:8081"
})

export async function getTags(): Promise<string[]> {

    const res = await api.get("/api/tags")

    return res.data

}

export async function getProblems(
    condition:SearchCondition
):Promise<Problem[]>{

    const res=await api.get("/api/problems",{

        params:{

            keyword:condition.keyword,

            tags:condition.tags.join(","),

            contestType:condition.contestType,

            minDifficulty:condition.minDifficulty ?? "",

            maxDifficulty:condition.maxDifficulty ?? ""

        }

    })

    return res.data

}