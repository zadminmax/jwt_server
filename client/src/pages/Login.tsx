import React, { SyntheticEvent, useState } from "react";
import {useNavigate} from "react-router-dom";
import cookie from 'react-cookies';

const Login=(props:{setName:(name:string)=>void})=>{
  
    const [user_name,setUsername]=useState('');
    const [password,setPassword]=useState('');
    const [redirect,setRedirect]=useState(false);
    let navigator=useNavigate();

    const submit=async (e:SyntheticEvent)=>{
      e.preventDefault();
      const response= await fetch('http://localhost:8000/api/login',{
          method:'POST',
          headers:{'Content-Type':'application/json'},
          credentials:'include',
          body:JSON.stringify({
            user_name,
            password
          })
      });

      const content=await response.json();
      if (content!=null&&content.err_code===0){
        cookie.save('token',content.data.token,{path:'/'})
        props.setName(content.data.user_name);
        setRedirect(true);
      }else{
        alert(content.msg);
        console.log(content.msg);
      }
  }
  if(redirect){
    navigator("/");
}
    return(
    <form onSubmit={submit}>
    <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
      <input className="form-control"  placeholder="Username" required
        onChange={e=>setUsername(e.target.value)}
      />
      <input type="password" className="form-control" placeholder="Password" required 
        onChange={e=>setPassword(e.target.value)}
      />
  
    <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
    
  </form>
    )
}

export default Login;