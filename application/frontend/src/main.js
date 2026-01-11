import "./main.scss";
import App from "./App.svelte";
import "@/assets/res/onLogsFont.css";
import { mount } from "svelte";

const app = mount(App, {
  target: document.getElementById("app"),
});

export default app;
