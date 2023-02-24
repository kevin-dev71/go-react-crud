// @ts-nocheck
import { useState } from 'react';

import './App.css';

function App() {
  const [users, setUsers] = useState('');
  return (
    <div className="container">
      <h1>Hello World</h1>
      <button
        onClick={async () => {
          const response = await fetch('http://localhost:4000/users');
          const data = await response.json();
          setUsers(data.data);
        }}
      >
        Get users
      </button>
      <div>{users}</div>
    </div>
  );
}

export default App;
