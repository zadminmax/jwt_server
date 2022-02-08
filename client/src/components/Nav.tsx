import React from "react";
import { Link } from "react-router-dom";
import cookie from 'react-cookies';

const Nav=(props:{name:string,setName:(name:string)=>void})=>{
  const logout=async ()=>{
    await fetch('http://localhost:8000/data/logout',{
                headers:{'Content-Type':'application/json','token':cookie.load("token")},
                credentials:'include',}
              )
     props.setName('');         
  }
  let menu;
  if (props.name===''){
    menu=(
      <ul className="navbar-nav me-auto mb-2 mb-md-0">
      <li className="nav-item">
        <Link className="nav-link" to="/login">Login</Link>
      </li>
      <li className="nav-item">
        <Link className="nav-link" to="/register">Register</Link>
      </li>
    </ul>
    )
  }else{
    menu=(
      <ul className="navbar-nav me-auto mb-2 mb-md-0">
      <li className="nav-item">
        <Link className="nav-link" to="/login" onClick={logout}>Logout</Link>
      </li>
    </ul>
    )
  }
    return(
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
        <div className="container-fluid">
          <Link className="navbar-brand" to="/">Home</Link>
          
          <div >
           {menu}
          </div>
        </div>
      </nav>
    )
}

export default Nav;