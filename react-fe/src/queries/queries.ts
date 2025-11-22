export const getDashboard = async () => {
    const res = await fetch('http://localhost:8080/dashboard');
    if (!res.ok) {
        throw new Error('Failed to fetch dashboard data');
    }
    return res.json();
}