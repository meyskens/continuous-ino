import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import registerServiceWorker from './registerServiceWorker';
import Navigation from "./components/navigation"
import { BrowserRouter as Router, Route } from 'react-router-dom';

import Builds from './components/builds';
import Build from './components/build';

ReactDOM.render(
    <div>
        <Router>
            <div>
                <Navigation />
                <div className="container">
                    <Route path="/builds" component={Builds} />
                    <Route path="/build/:id?" component={Build} />
                </div>
            </div>
        </Router>
    </div>, document.getElementById('root'));
registerServiceWorker();
