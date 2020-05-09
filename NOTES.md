# Notes on challenge

From the instructions it is unclear what should be assumed and what shouldn't,
for example:

  * should we assume that there is infrastructure already setup to run
    the artifiact created form this ie: k8's, ecs other?.

  * If not is terraform code wanted. If terraform code is wanted what assumptions
    if any should be made there? Is there a network setup an account setup,
    Iam persmissions, service discovery, load balancers security groups etc.


Assumptions I made:
  * The service executor of my chosing was already setup, in this example I used
    ECS.

  * CircleCI was the CI/CD of choice. Assumption being contexts were already setup
    for secrets the pipeline needed

  * Appropriate security baselines are already in place ie: seccomp / selinux / apparmour
    base profiles already instituted at the org

  * SSL conventions already exist. Since it wasn't clear on if the service should terminate
    its own ssl or if this should be considered for this.

  * Space and time complexity of the log parser.



There is circleci workflows and a terraform module both won't work due to some
underlying things not being there / values being set to "example" but both of them
are valid code.
