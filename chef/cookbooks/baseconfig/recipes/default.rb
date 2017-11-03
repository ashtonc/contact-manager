# Make sure the Apt package lists are up to date, so we're downloading versions that exist.
cookbook_file "apt-sources.list" do
  path "/etc/apt/sources.list"
end
execute 'apt_update' do
  command 'apt-get update'
end

# Base configuration recipe in Chef.
package "wget"
package "ntp"
cookbook_file "ntp.conf" do
  path "/etc/ntp.conf"
end
execute 'ntp_restart' do
  command 'service ntp restart'
end

# Go installation
package "golang"
execute 'get-mux' do
  environment 'GOPATH' => '/go'
  command 'go get -u github.com/gorilla/mux'
end

# Install tmux and start the server in the background
package "tmux"
execute 'create-server-session' do
  cwd '/vagrant'
  environment 'GOPATH' => '/go'
  command 'tmux new-session -d -s server'
end
execute 'start-server' do
  command "tmux send-keys -t server 'go run contactmanager.go' C-m"
end
