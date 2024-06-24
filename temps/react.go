package temps

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

func ReactFrame() {
	// Open the JSON file
	// file, err := os.Open("config.json")
	// if err != nil {
	// 	fmt.Println("Error opening JSON file:", err)
	// 	return
	// }
	// defer file.Close() // Defer closing the file until the function returns

	// // Decode the JSON content into the data structure
	var data Data
	data.BackTick = "`"
	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&data)
	// if err != nil {
	// 	fmt.Println("Error decoding JSON:", err)
	// 	return
	// }
	// #############################################################
	// #############################################################
	appjsx_tmpl, err := template.New("data").Parse(appTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	appjsx_file, err := os.Create("src/App.jsx")
	if err != nil {
		panic(err)
	}
	defer appjsx_file.Close()

	err = appjsx_tmpl.Execute(appjsx_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	mainjsx_tmpl, err := template.New("data").Parse(mainTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	mainjsx_file, err := os.Create("src/main.jsx")
	if err != nil {
		panic(err)
	}
	defer appjsx_file.Close()

	err = mainjsx_tmpl.Execute(mainjsx_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	indexcss_tmpl, err := template.New("data").Parse(indexTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	indexcss_file, err := os.Create("src/index.css")
	if err != nil {
		panic(err)
	}
	defer indexcss_file.Close()

	err = indexcss_tmpl.Execute(indexcss_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	inputjsx_tmpl, err := template.New("data").Parse(inputTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	inputjsx_file, err := os.Create("src/components/Input.jsx")
	if err != nil {
		panic(err)
	}
	defer inputjsx_file.Close()

	err = inputjsx_tmpl.Execute(inputjsx_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	iconjsx_tmpl, err := template.New("data").Parse(iconTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	iconjsx_file, err := os.Create("src/components/Icons.jsx")
	if err != nil {
		panic(err)
	}
	defer iconjsx_file.Close()

	err = iconjsx_tmpl.Execute(iconjsx_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	buttonjsx_tmpl, err := template.New("data").Parse(buttonTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	buttonjsx_file, err := os.Create("src/components/Button.jsx")
	if err != nil {
		panic(err)
	}
	defer buttonjsx_file.Close()

	err = buttonjsx_tmpl.Execute(buttonjsx_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	privatejsx_tmpl, err := template.New("data").Parse(privateTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	privatejsx_file, err := os.Create("src/components/PrivateRoute.jsx")
	if err != nil {
		panic(err)
	}
	defer privatejsx_file.Close()

	err = privatejsx_tmpl.Execute(privatejsx_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	homejsx_tmpl, err := template.New("data").Parse(homeTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	homejsx_file, err := os.Create("src/page/Home.jsx")
	if err != nil {
		panic(err)
	}
	defer homejsx_file.Close()

	err = homejsx_tmpl.Execute(homejsx_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	logjsx_tmpl, err := template.New("data").Parse(loginPageTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	logjsx_file, err := os.Create("src/page/Login.jsx")
	if err != nil {
		panic(err)
	}
	defer logjsx_file.Close()

	err = logjsx_tmpl.Execute(logjsx_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	storejsx_tmpl, err := template.New("data").Parse(loginStoreTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	storejsx_file, err := os.Create("src/store/login.js")
	if err != nil {
		panic(err)
	}
	defer storejsx_file.Close()

	err = storejsx_tmpl.Execute(storejsx_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	clientjsx_tmpl, err := template.New("data").Parse(clientTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	clientjsx_file, err := os.Create("src/store/client.js")
	if err != nil {
		panic(err)
	}
	defer clientjsx_file.Close()

	err = clientjsx_tmpl.Execute(clientjsx_file, data)
	if err != nil {
		panic(err)
	}

	// running go mod tidy finally
	if err := exec.Command("npm", "install").Run(); err != nil {
		fmt.Printf("error: %v \n", err)
	}

	if err := exec.Command("npx", "tailwindcss", "-i", "src/index.css", "-m", "-o", "src/main.css").Run(); err != nil {
		fmt.Printf("error: %v \n", err)
	}
}

var privateTemplate = `
import { Navigate, Outlet } from 'react-router-dom'
import { useLogInStore } from '../store/login' 

export function PrivateRoutes () {
   
    let authToken=useLogInStore((state)=> state.login_status)
    
        return (
            authToken ? <Outlet/> : <Navigate to='/login'/>
        )
}

`

var appTemplate = `
import "./main.css";
import { LogInPage } from "./page/Login";
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import { Toaster } from "react-hot-toast";
import { HomePage } from "./page/Home";
import { PrivateRoutes } from "./components/PrivateRoute";
import { useEffect } from "react";
import { useLogInStore } from "./store/login";
function App() {

  const loggedIn = useLogInStore((state)=> state.login_status)
  useEffect(() => {
  }, [loggedIn]);

  return (
    <section className="w-full h-full">
      <BrowserRouter>
        <div className="w-full h-full flex flex-row md:flex-col space-y-1 items-center overflow-hidden">
          <div className="w-full h-full flex flex-col justify-center bg-slate-50 shadow-xl  py-1 items-center  overflow-y-scroll scrollbar-none md:scrollbar-thin scrollbar-thumb-gray-300 scrollbar-track-gray-100">
              <Routes>
                <Route path="/login" element={<LogInPage />} />
                <Route element={<PrivateRoutes />}>
                  <Route path="/" element={<HomePage />} />
                  <Route path="/home" element={<HomePage />} />   
                </Route>
              </Routes>
              <Toaster />
          </div>
        </div>
      </BrowserRouter>
    </section>
  );
}

export default App;
`

var homeTemplate = `
import { useState } from "react"
import { TextInput, SingleInput } from "../components/Input";
import { NormalButton } from "../components/Button";
import { useLogInStore } from "../store/login";




export function HomePage(){ 
    const logout = useLogInStore((state) => state.resetTokenLogout);

    return (
        <div className="w-full h-full flex flex-col py-3" >
        <div className="mx-5 bg-slate-50 w-11/12 justify-center rounded-box">
            <img className=" 	aspect-ratio: auto; rounded-xl" src="/wing.svg" alt="wings" />
            <NormalButton label={"Logout"} handleClick={logout}/>
        </div>
        <div className="flex w-full flex-col">
        <div className="divider mx-10"></div> 
        </div>
        <div className="w-full h-full flex flex-col md:flex-row md:flex-wrap items-stretch pt-5 pb-5  px-2 md:px-10  space-x-1 md:space-x-7  lg:space-x-14 space-y-12 ">
            <div></div>
            <div className="card hover:z-50 mt-30 transition ease-in-out delay-150 hover:scale-95 w-11/12 md:w-5/12 lg:w-3/12 bg-base-100 shadow-xl">
            <figure className="-mt-5 rotate-0 mx-2" ><img className="rounded-xl" src="/blue.svg" alt="blue" /></figure>
            <div className="card-body">
                <h2 className="card-title">
                Admin Features
                <div className="badge bg-slate-200">
                </div>
                </h2>
                <p>
                    Register Users that are relevant to the App.
                </p>
                <div className="card-actions justify-end">
                <div className="badge badge-outline">UMS</div> 
                <div className="badge badge-outline">Blue Admin</div>
                </div>
            </div>
            </div>
            <div className="card hover:z-50 mt-30 transition ease-in-out delay-150 hover:scale-95 w-11/12 md:w-5/12 lg:w-3/12 bg-base-100 shadow-xl">
            <figure className="-mt-5 rotate-0 mx-2" ><img className="rounded-xl" src="/blue.svg" alt="blue" /></figure>
            <div className="card-body">
                <h2 className="card-title">
                Register Apps
                <div className="badge bg-slage-200"></div>
                </h2>
                <p>Register Apps that want to use this up manage based on roles</p>
                <div className="card-actions justify-end">
                <div className="badge badge-outline">UMS</div> 
                <div className="badge badge-outline">Blue Admin</div>
                </div>
            </div>
            </div>
            <div className="card hover:z-50 mt-30 transition ease-in-out delay-150 hover:scale-95 w-11/12 md:w-5/12 lg:w-3/12 bg-base-100 shadow-xl">
            <figure className="-mt-5 rotate-0 mx-2" ><img className="rounded-xl" src="/blue.svg" alt="blue" /></figure>
            <div className="card-body">
                <h2 className="card-title">
                Register Roles
                <div className="badge bg-slage-200"></div>
                </h2>
                <p>Register roles which can be part of an app at a time. Roles are a means used to group features. </p>
                <div className="card-actions justify-end">
                <div className="badge badge-outline">UMS</div> 
                <div className="badge badge-outline">Blue UMS</div>
                </div>
            </div>
            </div>
            <div className="card hover:z-50 mt-30 transition ease-in-out delay-150 hover:scale-95 w-11/12 md:w-5/12 lg:w-3/12 bg-base-100 shadow-xl">
            <figure className="-mt-5 rotate-0 mx-2" ><img className="rounded-xl" src="/blue.svg" alt="Shoes" /></figure>
            <div className="card-body">
                <h2 className="card-title">
                Register Features & Endpoints
                <div className="badge bg-slage-200"> </div>
                </h2>
                <p>Features are to be descriptive of what you want to acheive using the app. It can have servral endpoints assocatied with it.
                And Endpoint refres to the api "url" that is suppose to be called to access a given resource. An endpoint can only have one parent feature.     
                </p>
                <div className="card-actions justify-end">
                <div className="badge badge-outline">Blue UMS</div> 
                <div className="badge badge-outline">Role Based Admin</div>
                </div>
            </div>
            </div>
            <div className="card hover:z-50 mt-30 transition ease-in-out delay-150 hover:scale-95 w-11/12 md:w-5/12 lg:w-3/12 bg-base-100 shadow-xl">
            <figure className="-mt-5 rotate-0 mx-2" ><img className="rounded-xl" src="/blue.svg" alt="blue" /></figure>
            <div className="card-body">
                <h2 className="card-title">
                Blue Authentication Middleware
                <div className="badge bg-slage-200"></div>
                </h2>
                <p>An Implemented Middleware, where user will not be able to access endpoints if id does not have the required role for the endpoint.</p>
                <div className="card-actions justify-end">
                <div className="badge badge-outline">UMS</div> 
                <div className="badge badge-outline">Role Mangement</div>
                </div>
            </div>
            </div>
            <div className="card hover:z-50 mt-30 transition ease-in-out delay-150 hover:scale-95 w-11/12 md:w-5/12 lg:w-3/12 bg-base-100 shadow-xl">
            <figure className="-mt-5 rotate-0 mx-2" ><img className="rounded-xl" src="blue.svg" alt="blue" /></figure>
            <div className="card-body">
                <h2 className="card-title">
                RPC Endpoint Summary of Apps
                <div className="badge bg-slage-200"></div>
                </h2>
                <p>The RPC endpoints are for use to apps that are registerd.</p>
                <div className="card-actions justify-end">
                <div className="badge badge-outline">UMS</div> 
                <div className="badge badge-outline">RPC calls</div>
                </div>
            </div>
            </div>           
            
            
            
           
        </div>
        <div className="flex w-full flex-col">
        <div className="divider mx-10"></div> 
        </div>
        
        </div>
    )
}

`

var loginPageTemplate = `

import { useState } from "react";
import { Navigate } from "react-router-dom";
import { SingleInput, CheckBoxInput, PasswordInput } from "../components/Input";
import { NormalButton } from "../components/Button";
import { useLogInStore } from "../store/login";

export function LogInPage() {
  const login = useLogInStore((state) => state.setToken);
  const logged_in_state = useLogInStore((state) => state.login_status);

  const [show, setShow] = useState(false);
  const [values, setValues] = useState({
    grant_type: "authorization_code",
    email: "",
    password: "",
    token: "none",
  });

  async function handleClick() {
    login(values);
  }

  function handleShow() {
    setShow(!show);
  }

  const handleInputChange = (event) => {
    event.persist();
    const target = event.target;
    const value = target.type === "checkbox" ? target.checked : target.value;
    setValues((values) => ({
      ...values,
      [target.name]: value,
    }));
  };

  return logged_in_state ? (
    <Navigate to="/home" />
  ) : (
    <>
      <div className="items-center justify-center w-full flex h-auto p-2 m-2">
        <title>Login</title>
        <div className="flex w-full content-center items-center justify-center h-full">
          <div className="w-full sm:w-7/12 md:w-6/12  lg:w-5/12 xl:w-4/12 sm:px-4 py-5 bg-gray-200 rounded-xl shadow">
            <div className="relative bg-slate-100 opacity-90 -mt-12 mx-auto  flex flex-col min-w-0 break-words w-full mb-1 shadow-lg rounded-lg  border-0">
              <div className="flex-auto px-4 lg:px-10 py-10 pt-0">
                <div className="text-gray-400 text-center mb-3 font-bold">
                  <small>sign in with credentials</small>
                </div>
                <form>
                  <SingleInput
                    name="email"
                    label="Email"
                    inputType="text"
                    placeHolder="beimdegefu@gmail.com"
                    value={values.email}
                    handler={handleInputChange.bind(this)}
                  />

                  <PasswordInput
                    name="password"
                    label="Password"
                    inputType={show ? "text" : "password"}
                    placeHolder="password"
                    value={values.password}
                    handler={handleInputChange.bind(this)}
                  />

                  <br />
                  <CheckBoxInput
                    value={show}
                    label="Show Passowrd"
                    name="show"
                    handler={handleShow.bind(this)}
                  />
                  <br />

                  <NormalButton
                    label="Sign In"
                    handleClick={handleClick.bind(this)}
                  />
                </form>
              </div>
            </div>

            <div className="flex flex-wrap pt-0 relative">
              <div className="w-1/2 ">
                <a href="#pablo" className="text-black-200">
                  <small>Forgot password?</small>
                </a>
              </div>
              <div className="w-1/2 text-right">
                <a className="text-black-200" href="/signup">
                  <small>Create new account</small>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="w-full flex items-center justify-center pb-2 px-5">
        <div role="alert" className="alert">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            className="stroke-info shrink-0 w-6 h-6"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
            ></path>
          </svg>
          <span>
            To login as super user you can use superuser@mail.com. password is
            "default@123" for all users.
          </span>
        </div>
      </div>
    </>
  );
}
`

var mainTemplate = `
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)

`

var indexTemplate = `
@tailwind base;
@tailwind components;
@tailwind utilities;

`

var inputTemplate = `
import { useState } from "react"

export function SingleInput({ label, inputType, placeHolder,name, value, handler}){

    return(

        <div className="relative w-full mb-3">
            <label className="block uppercase text-gray-600 text-xs font-bold mb-2" htmlFor="grid-password">{label}</label>
            <input name={name} type={inputType} onChange={handler} aria-label={label} className="border-0 px-3 py-3 placeholder-gray-300 text-gray-600 bg-white rounded text-sm shadow focus:outline-none focus:ring w-full ease-linear transition-all duration-150" placeholder={placeHolder}  value={value}/>
        </div>

    )
}

export function ReadOnlySingleInput({ label, inputType, placeHolder,name, value, handler}){

    return(

        <div className="relative w-full mb-3">
            <label className="block uppercase text-gray-600 text-xs font-bold mb-2" htmlFor="grid-password">{label}</label>
            <input readonly name={name} type={inputType} onChange={handler} aria-label={label} className="border-0 px-3 py-3 placeholder-gray-300 text-gray-600 bg-white rounded text-sm shadow focus:outline-none focus:ring w-full ease-linear transition-all duration-150" placeholder={placeHolder}  value={value}/>
        </div>

    )
}

export function PasswordInput({ label, inputType, placeHolder,name, value, handler}){
    return(

        <div className="relative w-full mb-3">
            <label className="block uppercase text-gray-600 text-xs font-bold mb-2" htmlFor="grid-password">{label}</label>
            <input name={name} type={inputType} onChange={handler} aria-label={label} className="border-0 px-3 py-3 placeholder-gray-300 text-gray-600 bg-white rounded text-sm shadow focus:outline-none focus:ring w-full ease-linear transition-all duration-150" placeholder={placeHolder}  value={value}/>
        </div>

    )
}

export function TextInput({ label, placeHolder,name , value, handler,cn}){

    return(
        <div className="relative w-full mb-1">
            <label className="block uppercase  text-gray-600 text-xs font-bold mb-2" htmlFor="grid-password">{label}</label>
            <textarea name={name} onChange={handler} aria-label={label} className={{"{"}}{{.BackTick}}border-0 px-3 py-3 placeholder-gray-300 text-gray-600 bg-white rounded text-sm shadow focus:outline-none focus:ring w-full ease-linear transition-all ${cn ? cn : ""} duration-150{{.BackTick}}{{"}"}} placeholder={placeHolder} value={value} />
        </div>
        )
}

export function CheckBoxInput({label,name,handler, value}){
   
   return   (
   <div>
        <label className="inline-flex items-center cursor-pointer">
            <input name={name}  aria-label={label} value={value} checked={value} onChange={handler ? handler : null}  type="checkbox" className="form-checkbox border-0 rounded text-gray-700 ml-1 w-5 h-5 ease-linear transition-all duration-150" />
            <span className={label ? "ml-2 text-sm font-semibold text-gray-600" : "hidden"}>{label}</span>
        </label>
    </div>
    )
}

export function SelectInput({ data, name,handler, label}){
    
    return(
        <div className="relative w-full px-1 mb-3">
        <label className="block uppercase text-gray-600 text-xs font-bold mb-2" aria-label={label}>{label}</label>    
        <select name={name} className="block uppercase text-gray-600 text-xs font-bold w-full mb-2" onChange={handler}>
            <option defaultValue>
                select to add
            </option>
            {
            data.map((item,index)=>{
                return (
                    <option key={item.name+index} className="hover:bg-slate-300" value={item.id} >{item.name}</option>
                )
            })
            }
        </select>
        </div>
    )
}

export function SingleInputNoLabel({ label, inputType, placeHolder,name, value, handler}){

    return(

        <div className="relative w-full flex justify-center items-center ">
            <input name={name} type={inputType} onChange={handler} aria-label={label} className="border-0 resize-y px-3 py-3 placeholder-gray-300 text-gray-600 bg-white rounded text-sm shadow focus:outline-none focus:ring w-11/12 ease-linear transition-all duration-150" placeholder={placeHolder}  value={value}/>
        </div>

    )
}

export function TextInputNoLabel({ label, inputType, placeHolder,name, value, handler}){
    return(
        <div className="relative w-full flex justify-center items-center">
            <textarea name={name} type={inputType} onChange={handler} aria-label={label} className="border-0 resize-y px-3 py-3 placeholder-gray-300 text-gray-600 bg-white rounded text-sm shadow focus:outline-none focus:ring w-full ease-linear transition-all duration-150" placeholder={placeHolder}  value={value}/>
        </div>

    )
}

export function ReadOnlySingleInputNoLabel({ label, inputType, placeHolder,name, value}){

    return(

        <div className="relative w-full flex justify-center items-center">
            <input readOnly name={name} type={inputType} aria-label={label} className="border-0 px-3 py-3 placeholder-gray-300 text-gray-600 bg-white rounded text-sm shadow focus:outline-none focus:ring w-11/12 ease-linear transition-all duration-150" placeholder={placeHolder}  value={value}/>
        </div>

    )
}



`

var iconTemplate = `
export function TrashIcon(){
    return (
            <svg  className="stroke-red-400 rounded-xl w-6 h-6 transition ease-in-out duration-75 hover:-translate-y-1 hover:scale-105 hover:stroke-white hover:bg-red-400" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" >
                 <path strokeLinecap="round" strokeLinejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" />
            </svg>


            );
         
}

export function EditIcon(){
    return (
        <svg className="stroke-green-400 w-6 h-6 rounded-xl transition ease-in-out duration-75 hover:-translate-y-1 hover:scale-105 hover:stroke-white hover:bg-green-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" >
            <path strokeLinecap="round" strokeLinejoin="round" d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0115.75 21H5.25A2.25 2.25 0 013 18.75V8.25A2.25 2.25 0 015.25 6H10" />
        </svg>
      
            );
         
}




export function RefreshIcon(){
    return (
        <svg className="stroke-gray-300 w-6 h-6 transition rounded-xl ease-in-out duration-75 hover:-translate-y-1 hover:scale-105 hover:stroke-white hover:bg-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" >
          <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0l3.181 3.183a8.25 8.25 0 0013.803-3.7M4.031 9.865a8.25 8.25 0 0113.803-3.7l3.181 3.182m0-4.991v4.99" />
        </svg>
      
            );
         
}

export function TrayDownIcon(){
    return (
        <svg className="stroke-green-400 w-6 h-6 rounded-xl transition ease-in-out duration-75 hover:-translate-y-1 hover:scale-105 hover:stroke-white hover:bg-green-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" >
            <path strokeLinecap="round" strokeLinejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5M16.5 12L12 16.5m0 0L7.5 12m4.5 4.5V3" />
        </svg>
      
            );
         
}


export function TrayUpIcon(){
    return (
        <svg className="stroke-green-400 w-6 h-6 rounded-xl transition ease-in-out duration-75 hover:-translate-y-1 hover:scale-105 hover:stroke-white hover:bg-green-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" >
            <path strokeLinecap="round" strokeLinejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5m-13.5-9L12 3m0 0l4.5 4.5M12 3v13.5" />
        </svg>
      
            );
         
}

export function UpdateIcon(){
    return (
        <svg className="stroke-orange-300 rounded-xl w-6 h-6 transition ease-in-out duration-75 hover:-translate-y-1 hover:scale-105 hover:stroke-white hover:bg-orange-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" >
              <path strokeLinecap="round" strokeLinejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5m-13.5-9L12 3m0 0l4.5 4.5M12 3v13.5" />
        </svg>
      
            );
         
}


export function CancelIcon(){
    return (
        <svg className="stroke-green-300 rounded-xl w-6 h-6 transition ease-in-out duration-75 hover:-translate-y-1 hover:scale-105 hover:stroke-white hover:bg-green-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" >
              <path strokeLinecap="round" strokeLinejoin="round" d="M9.75 9.75l4.5 4.5m0-4.5l-4.5 4.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
      
            );
         
}

export function ThreeBarsIcon(){
    return (
        <svg className="stroke-gray-00 w-6 rounded-md h-6 transition ease-in-out duration-75 hover:-translate-y-1 hover:scale-105 hover:stroke-white hover:bg-gray-700" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" >
              <path strokeLinecap="round" strokeLinejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
        </svg>
      
            );


         
}


export function PlayIcon(){
    return (
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" className="w-6 h-6 stroke-green-400">
            <path strokeLinecap="round" strokeLinejoin="round" d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.347a1.125 1.125 0 0 1 0 1.972l-11.54 6.347a1.125 1.125 0 0 1-1.667-.986V5.653Z" />
        </svg>

    );
}

export function MuteIcon(){
    return (
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" className="w-6 h-6 stroke-gray-700 hover:stroke-gray-50 hover:bg-gray-700 bg-gray-300 rounded-2xl">
              <path strokeLinecap="round" strokeLinejoin="round" d="M17.25 9.75 19.5 12m0 0 2.25 2.25M19.5 12l2.25-2.25M19.5 12l-2.25 2.25m-10.5-6 4.72-4.72a.75.75 0 0 1 1.28.53v15.88a.75.75 0 0 1-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.507-1.938-1.354A9.009 9.009 0 0 1 2.25 12c0-.83.112-1.633.322-2.396C2.806 8.756 3.63 8.25 4.51 8.25H6.75Z" />

        </svg>

    );
}



export function VolumeIcon(){
    return (
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" className="w-6 h-6 stroke-gray-100">
              <path strokeLinecap="round" strokeLinejoin="round" d="M19.114 5.636a9 9 0 0 1 0 12.728M16.463 8.288a5.25 5.25 0 0 1 0 7.424M6.75 8.25l4.72-4.72a.75.75 0 0 1 1.28.53v15.88a.75.75 0 0 1-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.507-1.938-1.354A9.009 9.009 0 0 1 2.25 12c0-.83.112-1.633.322-2.396C2.806 8.756 3.63 8.25 4.51 8.25H6.75Z" />
        </svg>

    );
}


export function BackIcon(){
    return (
        <div className="rounded-xl bg-slate-300 hover:bg-slate-700">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6 stroke-gray-700 hover:stroke-gray-50">
            <path strokeLinecap="round" strokeLinejoin="round" d="m18.75 4.5-7.5 7.5 7.5 7.5m-6-15L5.25 12l7.5 7.5" />
            </svg>
        </div>

    )
}
`

var clientTemplate = `
import axios from "axios";
export const Client = axios.create({
  baseURL: "http://localhost:3500/api/v1",
  timeout: 10000,
});

`

var loginStoreTemplate = `
import { Client } from "./client";
import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";
import { jwtDecode } from "jwt-decode";

import toast from "react-hot-toast";

export const useLogInStore = create(
  persist(
    (set, get) => ({
      login_status: false,
      access_token: "anonymous",
      setToken: async (data) => {
	  	set((state) => ({
              ...state,
              login_status: true,
            }));

      },
      refToken: async () => {
        set((state) => ({
              ...state,
              login_status: true,
            }));
      },
      resetTokenLogout: () =>
        set({
          login_status: false,
        }),
    }),
    {
      name: "login-storage", // name of the item in the storage (must be unique)
      storage: createJSONStorage(() => sessionStorage), // (optional) by default, 'localStorage' is used
    }
  )
);

`

var buttonTemplate = `

export function NormalButton({ label, handleClick}){
    return (
      <button onClick={handleClick ? handleClick : null} aria-label={label} className="bg-gray-800 text-white active:bg-gray-600 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 my-2 w-full ease-linear transition-all duration-150" type="button">
        {label}
      </button>
    )
  }

export function TabNormalButton({ label, handleClick, token, index}){
    return (
      <button onClick={handleClick ? handleClick : null} aria-label={label} 
      className={token == index ?"bg-gray-400 text-black text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full ease-linear transition-all duration-150" 
        :
       "bg-gray-300 text-black text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full ease-linear transition-all duration-150" }
         type="button">
        {label}
      </button>
    )
  }
`
