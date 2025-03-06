import { useUserInfoStore } from '@/stores/userInfo';
import axios from 'axios';

const request = axios.create({
    baseURL: (import.meta.env.VITE_BACKEND_URL || '') + "/api/v1",
    timeout: import.meta.env.VITE_REQUEST_TIMEOUT,
    retry: 3,
});

request.interceptors.request.use(
    (config) => {
        // Set jwt header
        const user = useUserInfoStore();
        config.headers.Authorization = `Bearer ${user.jwt}`;
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

request.interceptors.response.use(
    (response) => {
        return response.data;
    },
    (error) => {
        return Promise.reject(error);
    }
);

export default request;