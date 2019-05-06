import React from 'react'

export default class App extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      isLoading: true
    }
  }
  componentDidMount() {
	WebAssembly.instantiateStreaming(fetch("http://localhost:3000"), go.importObject).then(async (result) => {
    go.run(result.instance)
    this.setState({ isLoading: false })

    const obj = generateObject()
    obj.onMessage = (mes) => {console.log("js callback" + mes)}
    console.log(obj.do())
	});
  }
  render() {
    return (
      this.state.isLoading ? 
      <div>Loading</div> :  
      <div>Go WASM React Example </div>
    )
  }
}