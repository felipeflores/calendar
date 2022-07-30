import React from 'react';

import { useRoutes } from 'react-router-dom';

import logo from './logo.svg';
import './App.css';

import ThemeRoutes from './routes/router';

function App() {
  const routing = useRoutes(ThemeRoutes);

  return (
      <div className="dark">{routing}</div>
  );
}

export default App;
