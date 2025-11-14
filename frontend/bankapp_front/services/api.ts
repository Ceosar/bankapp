import axios from 'axios';
import { API_BASE_URL } from '@/config/api';

const api = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});

export const getUsers = async () => {
    const response = await api.get('/users');
    return response.data.users;
};

export const createUser = async (userData: any) => {
    const response = await api.post('/users', userData);
    return response.data;
};

export default api;