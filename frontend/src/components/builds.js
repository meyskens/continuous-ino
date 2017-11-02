import React, { Component } from 'react';
import { Nav, NavItem, } from 'react-bootstrap';
import { LinkContainer } from 'react-router-bootstrap';
import BuildAttribute  from './buildattribute'

class Builds extends Component {
    constructor() {
        super();
        this.state = {
            builds: []
        };
    }

    componentDidMount() {
        fetch(`https://continuous-ino.eyskens.me/api/build/all`)
            .then(result => result.json())
            .then(builds => this.setState({ builds }))
    }

    render() {
        var buildsList = this.state.builds.map((build, key) => {
            return (
                <LinkContainer key={key} to={/build/ + build.id}>
                    <NavItem><BuildAttribute build={build}/> {build.id} - {build.repo}: {build.sha}</NavItem>
                </LinkContainer>
            );
        })


        return (
            <div>
                <h2>Builds</h2>
                <Nav bsStyle="pills" stacked activeKey={1}>
                    {buildsList}
                </Nav>
            </div>
        );
    }
}

export default Builds;
