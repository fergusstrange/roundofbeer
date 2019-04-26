import './App.css';

import React from 'react';
import axios from 'axios';

export default class App extends React.Component {
    state = {
        name: '',
    };

    handleSubmit = event => {
        event.preventDefault();
        axios.get(`/api/round`).then(res => this.setState({name: res.data.id}))
    };

    render() {
        return (
            <div>
                <form onSubmit={this.handleSubmit}>
                    <label>
                        Person Name:
                        <input type="text" name="name" value={this.state.name}/>
                    </label>
                    <button type="submit">Add</button>
                </form>
            </div>
        )
    }
}