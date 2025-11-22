export async function load({ fetch }) {
  const res = await fetch('http://localhost:8080/dashboard', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  });

  if (!res.ok) {
    throw new Error('Failed to load items');
  }

  const items = await res.json();

  return { items };
}
