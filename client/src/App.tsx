import { BrowserRouter, Redirect, Route, Switch } from 'react-router-dom';
import { Register } from "./Pages/Register/Register"
import { Login } from './Pages/Login/Login';
import { Home } from "./Pages/Home/Home"
import 'react-toastify/dist/ReactToastify.min.css';

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/" component={Login} exact />
        <Route path="/register" component={Register} />
        <Route path="/home"  render={(props) => {
          const component = sessionStorage.getItem('token') != null ? <Home></Home> : <Redirect to="/" />;

          return (
            <div>
              {component}
            </div>
          );
        }}/>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
