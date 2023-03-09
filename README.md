# Go-Svelte

Just some basic scaffolding for building a website with a go backend and a svelte frontend.

Svelte is set up to produce static pages which the backend can use as templates. The example uses gofibre for routing and templates. On page load the backend places page data in a global javascript variable named "payload" which can be accessed on the svelte onMount event to update the page.

First build the front-end, this creates the "static" directory
```bash
cd frontend
npm install
npm run build
cd ..
```

Then build and run the backend
```bash
go mod tidy
go build
./go-svelte
```

Visit http://localhost:8000 and you should see your page.

