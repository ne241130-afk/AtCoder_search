import axios from "axios"
import type { Problem } from "../types/problem"

const api = axios.create({
  baseURL: "http://localhost:8081",
})

export async function getProblems(
  keyword = "",
  tag = ""
): Promise<Problem[]> {

  const res = await api.get("/api/problems", {
    params: {
      keyword,
      tag,
    },
  })

  return res.data
}