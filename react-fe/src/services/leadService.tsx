import axios from 'axios';

const API_URL = '/lead/search';

export const searchLeads = async (query: string) => {
    try {
        const response = await axios.post(API_URL, { search: query });
        return response.data;
    } catch (error) {
        throw new Error('Error searching leads: ' + error);
    }
};