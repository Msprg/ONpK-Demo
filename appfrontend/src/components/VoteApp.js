import React, { Component } from 'react';
import Platform from './Platform';

class VoteApp extends Component {
  render () {    
    return (
      <main role="main">
        <div class="jumbotron">
          <div class="container">
            <h1 class="display-3">Platform Vote v2</h1>
            ❤ DevOps 2020 v2.10.3
          </div>
        </div>

        <div class="container">
          <div class="row">
            <div class="col-md-4">
              <Platform id="kubernetes" logo="kubernetes.png"/>
            </div>
            <div class="col-md-4">
              <Platform id="rancher" logo="rancher.png"/>
            </div>
            <div class="col-md-4">
              <Platform id="openshift" logo="openshift.png"/>
            </div>
          </div>
        </div>
      </main>
    )
  }
}

export default VoteApp;