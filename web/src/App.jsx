import React from 'react';
import { App } from 'antd';
import { HashRouter } from 'react-router-dom';
import { Routes } from '@/route/RouteRules.jsx';
import RouteGuard from '@/route/RouteGuard.jsx';

const MainApp = () => {
  return (
    <App>
      <HashRouter>
        <RouteGuard>
          <Routes />
        </RouteGuard>
      </HashRouter>
    </App>
  );
};

export default MainApp;
