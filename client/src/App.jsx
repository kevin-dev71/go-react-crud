// @ts-nocheck
import { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [name, setName] = useState('');
  const [users, setUsers] = useState([]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    await fetch(import.meta.env.VITE_API_URL + '/users', {
      method: 'POST',
      body: JSON.stringify({ name }),
      headers: {
        "Content-Type": "application/json"
      }
    });
    fetch(import.meta.env.VITE_API_URL + '/users').then(res => res.json()).then(data => setUsers(data?.data))
  };

  useEffect(() => {
    fetch(import.meta.env.VITE_API_URL + '/users').then(res => res.json()).then(data => setUsers(data?.data))
  }, [])
  

  return (
    <div className="container">
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="name"
          placeholder="Insert your name"
          onChange={(e) => setName(e.target.value)}
        />
        <button type="submit">Save</button>
      </form>

      {users.map(item => <div key={item._id}>{item.name}</div>)}
    </div>
  );
}

export default App;
