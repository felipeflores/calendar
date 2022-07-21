import React from 'react';
import { useRoutes } from 'react-router-dom';

import logo from './logo.svg';
import './App.css';

import ThemeRoutes from './routes/router';

function App() {
  const routing = useRoutes(ThemeRoutes);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
      <div className="dark">{routing}</div>
    </div>
  );
}

export default App;
