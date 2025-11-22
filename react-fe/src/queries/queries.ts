export const getDashboard = async () => {
    const res = await fetch('http://localhost:8080/dashboard');
    if (!res.ok) {
        throw new Error('Failed to fetch dashboard data');
    }
    return res.json();
}

export const search = async (query: string) => {
    const res = await fetch(`http://localhost:8080/lead/search?search=${query}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    });
    if (!res.ok) {
        throw new Error('Failed to search leads');
    }
    return res.json();
}