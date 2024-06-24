package temps

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

func CommonFrame() {
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
	// #####################################################################
	// Create the models directory if it does not exist
	err := os.MkdirAll("src", os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("src/page", os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("src/store", os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("src/utils", os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("src/components", os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("src/assets", os.ModePerm)
	if err != nil {
		panic(err)
	}

	// ############################################################
	git_tmpl, err := template.New("data").Parse(gitIgnoreTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	git_file, err := os.Create(".gitignore")
	if err != nil {
		panic(err)
	}
	defer git_file.Close()

	err = git_tmpl.Execute(git_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	server_tmpl, err := template.New("data").Parse(serverTemplate)
	if err != nil {
		panic(err)
	}

	//  using templates
	server_file, err := os.Create("server.js")
	if err != nil {
		panic(err)
	}
	defer server_file.Close()

	err = server_tmpl.Execute(server_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	docker_tmpl, err := template.New("data").Parse(dockerConfig)
	if err != nil {
		panic(err)
	}

	//  using templates
	docker_file, err := os.Create("ui.Dockerfile")
	if err != nil {
		panic(err)
	}
	defer docker_file.Close()

	err = docker_tmpl.Execute(docker_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	index_tmpl, err := template.New("data").Parse(indexConfig)
	if err != nil {
		panic(err)
	}

	//  using templates
	index_file, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer index_file.Close()

	err = index_tmpl.Execute(index_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	package_json_tmpl, err := template.New("data").Parse(packageJSON)
	if err != nil {
		panic(err)
	}

	//  using templates
	package_file, err := os.Create("package.json")
	if err != nil {
		panic(err)
	}
	defer package_file.Close()

	err = package_json_tmpl.Execute(package_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	postcss_tmpl, err := template.New("data").Parse(postCSSConfig)
	if err != nil {
		panic(err)
	}

	//  using templates
	postcss_file, err := os.Create("postcss.config.js")
	if err != nil {
		panic(err)
	}
	defer postcss_file.Close()

	err = postcss_tmpl.Execute(postcss_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	vite_tmpl, err := template.New("data").Parse(vitConfig)
	if err != nil {
		panic(err)
	}

	//  using templates
	vite_file, err := os.Create("vite.config.js")
	if err != nil {
		panic(err)
	}
	defer vite_file.Close()

	err = vite_tmpl.Execute(vite_file, data)
	if err != nil {
		panic(err)
	}

	// ############################################################
	tailwind_tmpl, err := template.New("data").Parse(tailwindConfig)
	if err != nil {
		panic(err)
	}

	//  using templates
	tailwind_file, err := os.Create("tailwind.config.js")
	if err != nil {
		panic(err)
	}
	defer tailwind_file.Close()

	err = tailwind_tmpl.Execute(tailwind_file, data)
	if err != nil {
		panic(err)
	}

	// running go mod tidy finally
	if err := exec.Command("npm", "install").Run(); err != nil {
		fmt.Printf("error: %v \n", err)
	}

}

var gitIgnoreTemplate = `
logs
*.log
npm-debug.log*
yarn-debug.log*
yarn-error.log*
pnpm-debug.log*
lerna-debug.log*

node_modules
dist
dist-ssr
*.local
.env
public
`

var packageJSON = `
{
  "name": "vite-zust-daisy",
  "private": true,
  "version": "0.0.0",
  "type": "module",
  "scripts": {
    "dev": "vite --host 0.0.0.0",
    "build": "vite build",
    "lint": "eslint . --ext js,jsx --report-unused-disable-directives --max-warnings 0",
    "preview": "vite preview --host 0.0.0.0",
    "serve": "vite build && bun server.js",
    "nom": "vite build && bun server.js"
  },
  "dependencies": {
    "@fortawesome/fontawesome-svg-core": "^6.5.1",
    "@tailwindcss/forms": "^0.5.7",
    "@tailwindcss/typography": "^0.5.10",
    "axios": "^1.6.2",
    "chart.js": "^3.9.1",
    "dashjs": "^4.7.2",
    "dotenv": "^16.3.1",
    "express": "^4.19.2",
    "jwt-decode": "^4.0.0",
    "media-chrome": "^2.1.0",
    "nodemon": "^3.1.0",
    "react": "^18.2.0",
    "react-chartjs-2": "^4.3.1",
    "react-dom": "^18.2.0",
    "react-hot-toast": "^2.4.1",
    "react-moment": "^1.1.3",
    "react-player": "^2.14.1",
    "react-router-dom": "^6.21.1",
    "react-video-js-player": "^1.1.1",
    "tailwind-scrollbar": "^3.0.5",
    "zustand": "^4.4.7"
  },
  "devDependencies": {
    "@types/react": "^18.2.43",
    "@types/react-dom": "^18.2.17",
    "@vitejs/plugin-react-swc": "^3.5.0",
    "autoprefixer": "^10.4.16",
    "daisyui": "^4.7.2",
    "eslint": "^8.55.0",
    "eslint-plugin-react": "^7.33.2",
    "eslint-plugin-react-hooks": "^4.6.0",
    "eslint-plugin-react-refresh": "^0.4.5",
    "postcss": "^8.4.32",
    "tailwindcss": "^3.4.0",
    "vite": "^5.0.8"
  }
}

`

var vitConfig = `
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
});

`

var tailwindConfig = `
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx}","*.{jsx,js,html}"],
  theme: {
    extend: {},
  },
  plugins: [
    require('tailwind-scrollbar')({ nocompatible: true }),
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    require("daisyui")
  
  ],
}
`

var indexConfig = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="/blue.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Vite+React+Zustuand+ DaisyUI</title>
  </head>
  <body class="font-sans">
    <main id="root" class="w-screen h-screen"></main>
    <script type="module" src="/src/main.jsx"></script>
  </body>
</html>

`

var postCSSConfig = `
export default {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
}


`

var dockerConfig = `
FROM oven/bun:alpine

RUN apk add npm 

RUN npm install -g npm@latest

WORKDIR /playground

COPY . .

COPY package.json ./
COPY package-lock.json ./

RUN npm install
RUN npm run build
RUN npm install express

COPY server.js ./

# Expose the default HAProxy port
EXPOSE 4173

CMD ["npm","run","serve"]
`

var serverTemplate = `
const express = require("express");
const path = require("path");

const app = express();
const port = 4173;

// Define the directory containing your static files
const staticFilesDirectory = path.join(__dirname, "dist");

// Serve static files from the 'public' directory
app.use(express.static(staticFilesDirectory));
// app.get("*", (req, res) => {
//   res.json({ path: sta });
// });
// Define a route handler for all GET requests other than the ones handled by static middleware
app.get("*", (req, res) => {
  // Send the index.html file for all non-static routes
  res.sendFile(path.join(staticFilesDirectory, "index.html"));
});

// Start the server
app.listen(port, "0.0.0.0", () => {
  console.log({{.BackTick}}Server is running on http://localhost:${port}{{.BackTick}});
});
`
