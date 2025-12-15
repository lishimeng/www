import { get } from "@/utils/request";

const BASE_URL = "/api/projects";

export const getProjectsAPi = () => get(BASE_URL);
