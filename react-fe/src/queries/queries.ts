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

export const createAssignMentMutation = async (newAssignment: {
    title: string;
    organization: string;
    contact: string;
    role: string;
    stack: string[];
    hourlyPrice: number;
    periodStartAt: string;
    periodEndAt: string;
    leadId: number;
    consultantId: number;
}) => {
    const res = await fetch('http://localhost:8080/assignment', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(newAssignment),
    });
    if (!res.ok) {
        throw new Error('Failed to search leads');
    }
    return res.json();
}
