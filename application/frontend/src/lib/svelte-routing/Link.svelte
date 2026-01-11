<script>
  import { getContext } from 'svelte';
  export let to = '';
  export let replace = false;
  const router = getContext('svelte-routing') || {};

  function onClick(e) {
    if (e.metaKey || e.ctrlKey || e.shiftKey || e.altKey) return;
    e.preventDefault();
    if (replace) history.replaceState({}, '', to);
    else history.pushState({}, '', to);
    window.dispatchEvent(new Event('popstate'));
  }
</script>

<a href={to} on:click|preventDefault={onClick}><slot /></a>
