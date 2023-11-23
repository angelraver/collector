<template>
  <div>
    <h1>Example: {{ name }}</h1>
    <form
      id="form"
      enctype="multipart/form-data"
      action="http://localhost:8001/imageadd"
      method="POST"
    >
      <input class="input file-input" type="file" name="file" multiple />
      <button class="button" type="submit">Submit</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'sample',
  props: {
    name: {
      type: String,
      required: false
    }
  },
  data() {
    return {
    }
  },
  mounted() {
    this.getUsers();
  },
  methods: {
    getUsers() {
      axios.get("http://localhost:8001/gameget")
      .then((response) => {
        console.log(response.data);
      })
      .catch((error) => {
        console.log(`The API returned an error: ${error}`);
      })
    }
  }
}
</script>
