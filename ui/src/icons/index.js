import Vue from 'vue'

import { library } from '@fortawesome/fontawesome-svg-core'

import {
  faHandMiddleFinger,
  faSearch,
  faList,
  faPlus,
  faPlusSquare,
  faCaretSquareUp,
  faCaretSquareDown,
  faTrash,
  faEdit
} from '@fortawesome/free-solid-svg-icons'

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(
  faHandMiddleFinger,
  faSearch,
  faList,
  faPlus,
  faPlusSquare,
  faCaretSquareUp,
  faCaretSquareDown,
  faTrash,
  faEdit
)

Vue.component('fontawesome', FontAwesomeIcon)

export default {}
