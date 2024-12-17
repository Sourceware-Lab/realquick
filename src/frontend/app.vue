<script setup>
const colorMode = useColorMode()

function toggleDarkMode() {
  colorMode.preference = colorMode.preference === 'dark' ? 'light' : 'dark'
}
</script>

<template>
  <div
    class="min-h-screen transition-colors duration-300" :class="[
      colorMode.preference === 'dark'
        ? 'bg-gradient-to-br from-gray-900 via-gray-800 to-indigo-900 text-white'
        : 'bg-gradient-to-br from-white via-indigo-50 to-blue-100 text-gray-900',
    ]"
  >
    <!-- Navigation Bar -->
    <nav
      class="backdrop-blur-lg border-b transition-colors duration-300" :class="[
        colorMode.preference === 'dark'
          ? 'bg-gray-900/50 border-gray-800'
          : 'bg-white/50 border-gray-200',
      ]"
    >
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <div class="flex-shrink-0 flex items-center">
              <div
                class="text-2xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-indigo-500 to-purple-500"
              >
                Real Quick
              </div>
            </div>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <a
                href="#"
                class="border-indigo-500 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium transition-all duration-300"
                :class="
                  colorMode.preference === 'dark'
                    ? 'text-white hover:text-indigo-400'
                    : 'text-gray-900 hover:text-indigo-600'
                "
              >
                Dashboard
              </a>
              <a
                href="#"
                class="border-transparent inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium transition-all duration-300"
                :class="
                  colorMode.preference === 'dark'
                    ? 'text-gray-300 hover:text-white hover:border-indigo-400'
                    : 'text-gray-600 hover:text-gray-900 hover:border-indigo-400'
                "
              >
                Schedule
              </a>
              <a
                href="#"
                class="border-transparent inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium transition-all duration-300"
                :class="
                  colorMode.preference === 'dark'
                    ? 'text-gray-300 hover:text-white hover:border-indigo-400'
                    : 'text-gray-600 hover:text-gray-900 hover:border-indigo-400'
                "
              >
                Tasks
              </a>
            </div>
          </div>
          <div class="flex items-center">
            <button
              class="p-2 rounded-lg transition-colors duration-300"
              :class="
                colorMode.preference === 'dark'
                  ? 'text-gray-300 hover:text-white'
                  : 'text-gray-600 hover:text-gray-900'
              "
              @click="toggleDarkMode"
            >
              <svg
                v-if="colorMode.preference === 'dark'"
                class="w-6 h-6"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"
                />
              </svg>
              <svg
                v-else
                class="w-6 h-6"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="relative">
      <!-- Decorative Elements -->
      <div class="absolute inset-0 overflow-hidden pointer-events-none">
        <div
          class="absolute -top-40 -right-40 w-80 h-80 bg-indigo-500 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-blob"
        />
        <div
          class="absolute -bottom-40 -left-40 w-80 h-80 bg-purple-500 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-blob animation-delay-2000"
        />
        <div
          class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-80 h-80 bg-pink-500 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-blob animation-delay-4000"
        />
      </div>

      <div class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8 relative">
        <!-- Hero Section -->
        <div
          class="backdrop-blur-lg overflow-hidden shadow-2xl rounded-2xl border transition-colors duration-300" :class="[
            colorMode.preference === 'dark'
              ? 'bg-gray-800/50 border-gray-700'
              : 'bg-white/50 border-gray-200',
          ]"
        >
          <div class="px-4 py-12 sm:px-6 relative">
            <div class="text-center">
              <h2
                class="text-4xl sm:text-5xl font-extrabold bg-clip-text text-transparent bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500 pb-2"
              >
                Manage Your College Life
              </h2>
              <p
                class="mt-4 max-w-2xl mx-auto text-xl"
                :class="colorMode.preference === 'dark' ? 'text-gray-300' : 'text-gray-600'"
              >
                Smart scheduling for smart students
              </p>
              <div class="mt-10 flex justify-center space-x-4">
                <button
                  class="px-8 py-3 rounded-lg bg-gradient-to-r from-indigo-600 to-purple-600 text-white font-medium transition-transform duration-300 hover:scale-105 hover:shadow-lg hover:shadow-indigo-500/25 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                  :class="
                    colorMode.preference === 'dark'
                      ? 'focus:ring-offset-gray-900'
                      : 'focus:ring-offset-white'
                  "
                >
                  Create Schedule
                </button>
                <button
                  class="px-8 py-3 rounded-lg font-medium border transition-transform duration-300 hover:scale-105 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2" :class="[
                    colorMode.preference === 'dark'
                      ? 'bg-gray-800 text-gray-300 border-gray-700 hover:text-white hover:border-indigo-500 focus:ring-offset-gray-900'
                      : 'bg-white text-gray-600 border-gray-200 hover:text-gray-900 hover:border-indigo-500 focus:ring-offset-white',
                  ]"
                >
                  View Demo
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Feature Grid -->
        <div class="mt-16">
          <div class="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
            <!-- Feature 1 -->
            <div
              class="group backdrop-blur-lg overflow-hidden rounded-xl border transition-all duration-300 hover:border-indigo-500 hover:shadow-lg hover:shadow-indigo-500/25" :class="[
                colorMode.preference === 'dark'
                  ? 'bg-gray-800/50 border-gray-700'
                  : 'bg-white/50 border-gray-200',
              ]"
            >
              <div class="p-6">
                <div
                  class="w-12 h-12 rounded-lg bg-gradient-to-br from-indigo-500 to-purple-500 flex items-center justify-center mb-4"
                >
                  <svg
                    class="w-6 h-6 text-white"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                    />
                  </svg>
                </div>
                <h3
                  class="text-xl font-semibold mb-2"
                  :class="colorMode.preference === 'dark' ? 'text-white' : 'text-gray-900'"
                >
                  Smart Scheduling
                </h3>
                <p :class="colorMode.preference === 'dark' ? 'text-gray-400' : 'text-gray-600'">
                  Automatically organize your classes, study sessions, and
                  activities
                  <!-- This is a commit of the message for the git bot -->
                </p>
              </div>
            </div>

            <!-- Feature 2 -->
            <div
              class="group backdrop-blur-lg overflow-hidden rounded-xl border transition-all duration-300 hover:border-indigo-500 hover:shadow-lg hover:shadow-indigo-500/25" :class="[
                colorMode.preference === 'dark'
                  ? 'bg-gray-800/50 border-gray-700'
                  : 'bg-white/50 border-gray-200',
              ]"
            >
              <div class="p-6">
                <div
                  class="w-12 h-12 rounded-lg bg-gradient-to-br from-purple-500 to-pink-500 flex items-center justify-center mb-4"
                >
                  <svg
                    class="w-6 h-6 text-white"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"
                    />
                  </svg>
                </div>
                <h3
                  class="text-xl font-semibold mb-2"
                  :class="colorMode.preference === 'dark' ? 'text-white' : 'text-gray-900'"
                >
                  Task Management
                </h3>
                <p :class="colorMode.preference === 'dark' ? 'text-gray-400' : 'text-gray-600'">
                  Keep track of assignments, deadlines, and study goals
                </p>
              </div>
            </div>

            <!-- Feature 3 -->
            <div
              class="group backdrop-blur-lg overflow-hidden rounded-xl border transition-all duration-300 hover:border-indigo-500 hover:shadow-lg hover:shadow-indigo-500/25" :class="[
                colorMode.preference === 'dark'
                  ? 'bg-gray-800/50 border-gray-700'
                  : 'bg-white/50 border-gray-200',
              ]"
            >
              <div class="p-6">
                <div
                  class="w-12 h-12 rounded-lg bg-gradient-to-br from-pink-500 to-indigo-500 flex items-center justify-center mb-4"
                >
                  <svg
                    class="w-6 h-6 text-white"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
                    />
                  </svg>
                </div>
                <h3
                  class="text-xl font-semibold mb-2"
                  :class="colorMode.preference === 'dark' ? 'text-white' : 'text-gray-900'"
                >
                  Smart Reminders
                </h3>
                <p :class="colorMode.preference === 'dark' ? 'text-gray-400' : 'text-gray-600'">
                  Never miss a class or assignment with timely notifications
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <footer
      class="backdrop-blur-lg border-t mt-12 transition-colors duration-300" :class="[
        colorMode.preference === 'dark'
          ? 'bg-gray-900/50 border-gray-800'
          : 'bg-white/50 border-gray-200',
      ]"
    >
      <div
        class="max-w-7xl mx-auto py-12 px-4 sm:px-6 md:flex md:items-center md:justify-between lg:px-8"
      >
        <div class="flex justify-center space-x-6 md:order-2">
          <a
            href="#"
            class="transition-colors duration-300" :class="[
              colorMode.preference === 'dark'
                ? 'text-gray-400 hover:text-indigo-400'
                : 'text-gray-600 hover:text-indigo-600',
            ]"
          >
            <span class="sr-only">GitHub</span>
            <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
              <path
                fill-rule="evenodd"
                d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z"
                clip-rule="evenodd"
              />
            </svg>
          </a>
        </div>
        <div class="mt-8 md:mt-0 md:order-1">
          <p
            class="text-center text-base" :class="[
              colorMode.preference === 'dark' ? 'text-gray-400' : 'text-gray-600',
            ]"
          >
            &copy; 2024 Real Quick. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>

<style>
.animate-blob {
  animation: blob 7s infinite;
}

.animation-delay-2000 {
  animation-delay: 2s;
}

.animation-delay-4000 {
  animation-delay: 4s;
}

@keyframes blob {
  0% {
    transform: translate(0px, 0px) scale(1);
  }
  33% {
    transform: translate(30px, -50px) scale(1.1);
  }
  66% {
    transform: translate(-20px, 20px) scale(0.9);
  }
  100% {
    transform: translate(0px, 0px) scale(1);
  }
}
</style>
