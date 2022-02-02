<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>hello!!!</h1>
        <sample name="Pepe"></sample>
        <img src="../public/pics/1643688275619806000.png" >
      </div>
    </div>
  </div>
</template>

<script>
// import axios from 'axios';
import sample from './components/sample.vue'

export default {
  components: { sample },
  name: 'App',

  data() { return {
    websiteUrl: '',
    thumbnailUrl: '',
  } },

  methods: {
    // makeWebsiteThumbnail() {
    //   axios.post("http://localhost:8001/api/thumbnail", {
    //     url: this.websiteUrl,
    //   })
    //   .then((response) => {
    //     this.thumbnailUrl = response.data.screenshot;
    //   })
    //   .catch((error) => {
    //     window.alert(`The API returned an error: ${error}`);
    //   })
    // }
  }
}
</script>
