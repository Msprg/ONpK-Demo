import React, { Component } from 'react';
import axios from 'axios'

class Vote extends Component{
  constructor (props) {
    super(props);

    //configure the APIHOSTPORT, this is the public IP address of the host that the API server is running on
    this.APIHOSTPORT = `${window._env_.REACT_APP_APIHOSTPORT}`;

    this.state = {
      vote: this.props.count
    }
    
    this.handleClick = this.handleClick.bind(this)
  }

  handleClick () {
    //implement the handleClick function which will be called when the user clicks on the voting button
    //this invokes an AJAX request to the API to vote on the current platform
    var url = `http://${this.APIHOSTPORT}/platforms/${this.props.id}/vote`;
    axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';
    axios.get(url)
      .then(response => this.setState({vote: this.state.vote+1}))
  }

  render () {
    //provide implementation for the render function to render the HTML for the Vote component
    return (
      <div id={this.props.id}>
        <button className='button' onClick={this.handleClick} type="button" class="btn btn-outline-success">+1</button>
        <div>
          <b>Votes</b>: {this.state.vote}
        </div>
      </div>
    )
  }
}

//export the Vote class, allows the Platform component to import it
export default Vote;