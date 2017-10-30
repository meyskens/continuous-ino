import React, { Component } from 'react';
import { Nav, NavItem, } from 'react-bootstrap';

class Builds extends Component {
    handleSelect(selectedKey) {
        alert(`selected ${selectedKey}`);
    }
    render() {
        return (
            <div>
                <h2>Builds</h2>
                <Nav bsStyle="pills" stacked activeKey={1} onSelect={this.handleSelect}>
                    <NavItem eventKey={1} href="/home">NavItem 1 content</NavItem>
                    <NavItem eventKey={2} title="Item">NavItem 2 content</NavItem>
                    <NavItem eventKey={3} disabled>NavItem 3 content</NavItem>
                </Nav>
            </div>
        );
    }
}

export default Builds;
