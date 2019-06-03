<template>
  <div class="container">
    <b-tabs vertical class="tc-tabs-v">
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
    addReplicants: function (values) {
      if (this.debug) {
        console.debug('adding replicants', values)
      }
      values.forEach(v => {
        this.replicants.push(v)
      })
      this.displayed = this.replicants
    },
    addReplicant: function (value) {
      if (this.debug) {
        console.debug('adding replicant', value)
      }
      this.replicants.push(value)
      this.displayed = this.replicants
    },
    removeReplicant: function (value) {
      if (this.debug) {
        console.debug('removing replicant', value)
      }
      this.replicants = this.replicants.filter(v => v.id !== value.id)
      this.displayed = this.replicants
    }
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
