import React, { SyntheticEvent, useState } from "react";
import {useNavigate} from "react-router-dom";


const Register=()=>{
    const [user_name,setUsername]=useState('');
    const [name,setName]=useState('');
    const [password,setPassword]=useState('');
    const [email,setEmail]=useState('');
    const [redirect,setRedirect]=useState(false);
    let navigator=useNavigate();

    const submit=async (e:SyntheticEvent)=>{
        e.preventDefault();
        const response= await fetch('http://localhost:8000/api/register',{
            method:'POST',
            headers:{'Content-Type':'application/json'},
            body:JSON.stringify({
               user_name,
                name,
                email,
                password
            })
        });

        const content=await response.json();
        if(content!=null&&content.err_code===0){
          setRedirect(true);
        }else{
          alert(content.msg);
          console.log(content.msg);
        }
        
    }

    if(redirect){
        navigator("/login");
    }

    return(
        <form onSubmit={submit}>
        <h1 className="h3 mb-3 fw-normal">Please register</h1>
          <input className="form-control"  placeholder="UserName" required 
            onChange={e=>setUsername(e.target.value)}
          />
          <input className="form-control"  placeholder="RealName"  
            onChange={e=>setName(e.target.value)}
          />

          <input type="password" className="form-control" placeholder="Password" required 
           onChange={e=>setPassword(e.target.value)}
          />

          <input type="email" className="form-control"  placeholder="name@example.com" required 
            onChange={e=>setEmail(e.target.value)}
          />
      
        <button className="w-100 btn btn-lg btn-primary" type="submit">Register</button>
        
      </form>
        )
}

export default Register;