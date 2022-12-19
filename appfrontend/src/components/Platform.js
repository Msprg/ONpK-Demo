import React, { Component } from 'react';
import Vote from './Vote';
import axios from 'axios';

class Platform extends Component {
  constructor () {
    super();

    //configure the APIHOSTPORT, this is the public IP address of the host that the API server is running on
    this.APIHOSTPORT = `${window._env_.REACT_APP_APIHOSTPORT}`;

    this.state = {
      platform: {},
      loaded: false
    }
  }

  componentDidMount () {
    //provide implementation to request platform details for current platform from the API server
    axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';
    axios.get(`http://${this.APIHOSTPORT}/platforms/${this.props.id}`).then(
      response => this.setState({
        platform: response.data,
        loaded: true
      })
    );
  }

  render () {
    //provide implementation for the render function to render the HTML for the Platform component
    if (this.state.loaded) {
      var usecase = this.state.platform.codedetail.usecase;
      var rank = this.state.platform.codedetail.rank;
      var homepage = this.state.platform.codedetail.homepage;
      var votecount = this.state.platform.codedetail.votes;

      return (
        <div class="container">
          <h2>{this.props.name}</h2>
          <p><Vote id={this.props.id} count={votecount}/></p>

          <p><b>Uses</b>: {usecase}</p>
          <p><b>Rank</b>: {rank}</p>
          <p><b>Homepage</b>: {homepage}</p>
          <div class="container">
            <div class="row">
              <div class="col">
                <div class="parent">
                  <img src={"./img/" + this.props.logo} alt="logo" class="center-block"/>
                </div>
              </div>
            </div>
          </div>
        </div>
      )
    }
    return <div></div>;
  }
}

export default Platform;