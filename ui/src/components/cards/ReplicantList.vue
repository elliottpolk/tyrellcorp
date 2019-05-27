<template>
  <b-card header-tag="header" bg-variant="light" class="mt-3 w-75">
    <template slot="header">
      <div class="d-flex flex-row font-weight-bold pt-0">Replicants</div>
    </template>

    <b-row></b-row>

    <b-row>
      <b-col>
        <b-table
          striped
          hover
          fixed
          bordered
          :items="displayed"
          :fields="fields"
          :sort-compare="sortCompare"
          :sort-by.sync="sortBy"
          :sort-desc.sync="sortDesc"
        >
          <template slot="name" slot-scope="row">
            <div :title="row.item.id">{{ row.item.name }}</div>
          </template>

          <template slot="show_details" slot-scope="row">
            <b-button variant="link" size="sm" @click="row.toggleDetails">
              <fontawesome :icon="row.detailsShowing ? 'caret-square-up' : 'caret-square-down'"/>
            </b-button>
          </template>

          <template slot="row-details" slot-scope="row">
            <b-card no-body class="d-flex flex-row">{{ row.item }}</b-card>
          </template>
        </b-table>
      </b-col>
    </b-row>
  </b-card>
</template>

<script>

export default {
  props: {
    parentState: Object
  },

  data () {
    return {
      fields: [
        { key: 'name', sortable: true },
        { key: 'package', sortable: true },
        { key: 'state', sortable: true },
        { key: 'show_details', label: 'Details', sortable: false }
      ],
      replicants: [],
      displayed: [],
      sortBy: '',
      sortDesc: false,
      api: {
        baseUrl: '/api/v1/specs'
      }
    }
  },

  mounted () {

  },

  created () {
    this.$http.get(this.api.baseUrl, { request_id: 'something_random' }).then(res => {
      let replicants = []
      res.body.payload.forEach(el => {
        replicants.push(el)
      })
      this.displayed = this.replicants = replicants
    }).catch(err => {
      console.error(err)
    })
  },

  watch: {
    replicants: function (val) {
      console.debug('changed')
      this.parentState.replicants = this.replicants
    },
    parentState: function (val) {
      console.debug('parent changed')
    }
  },

  methods: {
    sortCompare (a, b, key) {
      // handle nested keys
      key.split('.').forEach(k => {
        a = a[k]
        b = b[k]
      })

      if (!a || a === null) {
        return 1
      }

      if (!b || b === null) {
        return 0
      }

      if (typeof a === 'number' && typeof b === 'number') {
        // if both compared fields are native numbers
        return a < b ? -1 : (a > b ? 1 : 0)
      } else {
        // stringify the field data and use String.localeCompare
        return a.localCompare(b, 'en', { sensitivity: 'base' })
      }
    }
  }
}
</script>

<style>
</style>
