<script>
  import { setContext, onMount, onDestroy } from 'svelte';
  import { writable } from 'svelte/store';
  export let url = '';
  export let basepath = '';

  const initial = url && url.length > 0 ? url : window && window.location ? window.location.pathname : '/';
  const location = writable({ pathname: initial });

  setContext('svelte-routing', { location, basepath });

  function onPop() {
    location.set({ pathname: window.location.pathname });
  }

  onMount(() => {
    window.addEventListener('popstate', onPop);
  });

  onDestroy(() => {
    window.removeEventListener('popstate', onPop);
  });
</script>

<slot />
