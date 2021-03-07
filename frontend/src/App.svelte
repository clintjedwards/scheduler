<script>
  import { link, Route, Router } from "svelte-routing";
  import { globalHistory } from "svelte-routing/src/history";
  import AddEmployee from "./components/AddEmployee.svelte";
  import AddPosition from "./components/AddPosition.svelte";
  import AddProgram from "./components/AddProgram.svelte";
  import Employees from "./components/Employees.svelte";
  import Footer from "./components/Footer.svelte";
  import ManageEmployee from "./components/ManageEmployee.svelte";
  import ManagePosition from "./components/ManagePosition.svelte";
  import ManageProgram from "./components/ManageProgram.svelte";
  import Positions from "./components/Positions.svelte";
  import Programs from "./components/Programs.svelte";
  import Schedules from "./components/Schedules.svelte";

  let pathname = window.location.pathname;

  globalHistory.listen(({ location, _ }) => {
    pathname = location.pathname;
  });

  let routes = [
    { path: "/employees", text: "Employees" },
    { path: "/positions", text: "Positions" },
    { path: "/programs", text: "Programs" },
    { path: "/schedules", text: "Schedules" },
  ];
</script>

<main class="mx-auto w-3/4 flex flex-col font-body">
  <div id="content" class="flex-1">
    <img
      class="mx-auto"
      id="logo"
      src="/images/schedule.svg"
      alt="Scheduler Logo"
    />
    <h1 class="font-heading text-6xl text-center text-orange mb-10">
      Scheduler
    </h1>
    <Router>
      <nav class="text-2xl font-sm-heading text-center mb-10">
        {#each routes as route}
          <a
            class:selected={route.path === pathname}
            class="mr-10 text-gray-500 ease-in-out transition duration-700"
            href={route.path}
            use:link
          >
            {route.text}
          </a>
        {/each}
      </nav>
      <div>
        <Route path="employees" component={Employees} />
        <Route path="employees/add" component={AddEmployee} />
        <Route path="employees/:id" component={ManageEmployee} />
        <Route path="positions" component={Positions} />
        <Route path="positions/add" component={AddPosition} />
        <Route path="positions/:id" component={ManagePosition} />
        <Route path="programs" component={Programs} />
        <Route path="programs/add" component={AddProgram} />
        <Route path="programs/:id" component={ManageProgram} />
        <Route path="schedules" component={Schedules} />
        <Route path="schedules/add" component={Schedules} />
        <Route path="schedules/:id" component={Schedules} />
      </div>
    </Router>
  </div>
  <div>
    <Footer />
  </div>
</main>

<style>
  :global(html) {
    overflow-y: scroll;
  }

  #logo {
    max-width: 180px;
  }

  main {
    min-height: 95vh;
  }

  .selected {
    border-bottom: 1px solid #ff3e00;
    color: black;
  }
</style>
