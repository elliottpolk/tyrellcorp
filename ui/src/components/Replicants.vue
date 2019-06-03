<template>
  <div class="container">
    <b-tabs vertical v-model="sharedState.tabIndex" class="tc-tabs-v">
      <b-tab active>
        <template slot="title">
          <fontawesome icon="list"/>
        </template>
        <replicant-list :parentState="sharedState"/>
      </b-tab>
      <b-tab>
        <template slot="title">
          <fontawesome icon="plus"/>
        </template>
        <add-replicant :parentState="sharedState"/>
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>
import ReplicantList from '@/components/cards/ReplicantList'
import ReplicantAdd from '@/components/cards/ReplicantAdd'

let store = {
  debug: true,
  state: {
    replicants: [],
    tabIndex: 0
  }
}

export default {
  data () {
    return {
      sharedState: store.state
    }
  },

  created () {

  },

  methods: {

  },

  components: {
    'replicant-list': ReplicantList,
    'add-replicant': ReplicantAdd
  }
}
</script>

<style>
</style>
