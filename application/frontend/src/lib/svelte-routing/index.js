import Router from './Router.svelte';
import Route from './Route.svelte';
import Link from './Link.svelte';

export function navigate(to, { replace = false } = {}) {
  if (replace) history.replaceState({}, '', to);
  else history.pushState({}, '', to);
  window.dispatchEvent(new Event('popstate'));
}

export { Router, Route, Link };

export default {
  Router,
  Route,
  Link,
  navigate,
};
